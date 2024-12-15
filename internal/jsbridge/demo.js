const cp = require('child_process')


async function main() {
  const p = cp.spawn('./dist/jsbridge', {
    stdio: ['pipe', 'inherit', 'inherit', 'ipc'],
  })
  
  p.on('message', (msg) => {
    console.log(msg)
  })

  p.on('error', (e) => {
    console.log(e)
  })

  while (true) {
    if (typeof p.exitCode === 'number') {
      break
    }
    p.send("ping")
    await new Promise((res) => setTimeout(res, 1e3))
  }
}

main()

