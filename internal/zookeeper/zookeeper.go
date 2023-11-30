package zookeeper

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"sort"
	"sync/atomic"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
)

const ELECTION_NODE = "/election"

type createMode int

var (
	createModePersistent          createMode = 0
	createModeEphemeral           createMode = zk.FlagEphemeral
	createModeEphemeralSequential createMode = zk.FlagEphemeral | zk.FlagSequence
)

var errNodeNotExists = errors.New("node not exists")
var errNotAChildOfParent = errors.New("not a child of parent")

type ZookeeperService struct {
	conn     *zk.Conn
	username string
	password string
	IsMaster atomic.Bool
}

func NewZookeeper(config config.AppConfig) (*ZookeeperService, error) {
	conn, _, err := zk.Connect([]string{fmt.Sprintf("%s:%s", config.ZKHost, config.ZKPort)}, time.Minute*30)
	if err != nil {
		return nil, errgo.Wrap(err, "zk.Connect")
	}
	s := &ZookeeperService{
		conn:     conn,
		username: config.ZKUserName,
		password: config.ZKPassword,
	}
	err = s.createAllParentNodes()
	if err != nil {
		return nil, err
	}
	err = createNodeInElectionZnode(s, fmt.Sprintf("%s:%d", config.HTTPHost, config.HTTPPort))
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *ZookeeperService) exists(p string) (bool, error) {
	ok, _, err := s.conn.Exists(ELECTION_NODE)
	if err != nil {
		return false, errgo.Wrap(err, fmt.Sprintf("conn.Exists %s", p))
	}
	return ok, nil
}

func createZnode[T any](s *ZookeeperService, p string, mode createMode, data T) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", errgo.Wrap(err, "json.Marshal")
	}
	acl := zk.WorldACL(zk.PermAll)
	path, err := s.conn.Create(p, bytes, int32(mode), acl)
	if err != nil {
		return "", errgo.Wrap(err, fmt.Sprintf("conn.Create %s", p))
	}
	return path, nil
}

func (s *ZookeeperService) createAllParentNodes() error {
	ok, err := s.exists(ELECTION_NODE)
	if err != nil {
		return err
	}
	if !ok {
		_, err = createZnode(s, ELECTION_NODE, createModePersistent, "election node")
		if err != nil {
			return err
		}
	}
	return nil
}

func readZNodeData[T any](s *ZookeeperService, p string, out *T) error {
	data, _, err := s.conn.Get(p)
	if err != nil {
		return errgo.Wrap(err, "conn.Get")
	}
	err = json.Unmarshal(data, out)
	if err != nil {
		return errgo.Wrap(err, "json.Unmarshal")
	}
	return nil
}

func (s *ZookeeperService) getChildren(path string) ([]string, error) {
	children, _, err := s.conn.Children(ELECTION_NODE)
	if err != nil {
		return nil, errgo.Wrap(err, "conn.Children")
	}
	return children, nil
}

func (s *ZookeeperService) getPrevSiblingNode(parentPath string, refPath string) (string, error) {
	name := path.Base(refPath)
	children, err := s.getChildren(parentPath)
	if err != nil {
		return "", err
	}
	sort.Strings(children)
	i := sort.SearchStrings(children, name)
	if i == len(children) {
		return "", errNotAChildOfParent
	}
	if i > 0 {
		return path.Join(parentPath, children[i-1]), nil
	} else {
		return "", nil
	}
}

func GetLeaderNodeData[T any](s *ZookeeperService, out *T) error {
	ok, err := s.exists(ELECTION_NODE)
	if err != nil {
		return err
	}
	if !ok {
		return errNodeNotExists
	}
	children, err := s.getChildren(ELECTION_NODE)
	if err != nil {
		return err
	}
	if len(children) == 0 {
		return nil
	}
	sort.Strings(children)
	err = readZNodeData[T](s, path.Join(ELECTION_NODE, children[0]), out)
	if err != nil {
		return err
	}
	return nil
}

func createNodeInElectionZnode[T any](s *ZookeeperService, data T) error {
	ok, err := s.exists(ELECTION_NODE)
	if err != nil {
		return err
	}
	if !ok {
		_, err = createZnode(s, ELECTION_NODE, createModePersistent, "election node")
		if err != nil {
			return err
		}
	}
	path, err := createZnode(s, path.Join(ELECTION_NODE, "node"), createModeEphemeralSequential, data)
	if err != nil {
		return err
	}
	return s.tryWatchPrevSibling(ELECTION_NODE, path)
}

func (s *ZookeeperService) tryWatchPrevSibling(parentPath string, refPath string) error {
	prevSibling, err := s.getPrevSiblingNode(ELECTION_NODE, refPath)
	if err != nil {
		return err
	}
	if len(prevSibling) == 0 {
		s.IsMaster.Store(true)
		return nil
	} else {
		ok, _, eventsChan, err := s.conn.ExistsW(prevSibling)
		if err != nil {
			return errgo.Wrap(err, "conn.ExistsW")
		}
		if ok {
			go func() {
				for msg := range eventsChan {
					if msg.Type == zk.EventNodeDeleted {
						s.tryWatchPrevSibling(parentPath, refPath)
						break
					}
				}
			}()
			return nil
		} else {
			return s.tryWatchPrevSibling(parentPath, refPath)
		}
	}
}
