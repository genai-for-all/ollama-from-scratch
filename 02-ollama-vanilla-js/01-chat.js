//let ollamaUrl = "http://bob.local:11434"
let ollamaUrl = "http://host.docker.internal:11434"
let model = "deepseek-coder"

/*
  with: deepseek-coder
  temperature: 0.5
  predictRepeatLastN: 2

  role: the role of the message, either system, user or assistant

  - system for the instructions
  - assistant for the history of the answer
  - user for the human 

*/

async function postData(url = "", data = {}) {
  const response = await fetch(url, {
    method: "POST",
    cache: "no-cache",
    headers: {
      "Content-Type": "application/json; charset=utf-8",
    },
    body: JSON.stringify(data)
  })
  return response.json()
}

let systemContent = `You are an expert in computer programming.
Please make friendly answer for the noobs.
Add source code examples if you can.`

let userContent = `I need a clear explanation regarding the following question:
Can you create a "hello world" program in Golang?
And, please, be structured with bullet points`

let messages = [
  {role:"system", content: systemContent},
  {role:"user", content: userContent}
]

postData(`${ollamaUrl}/api/chat`, {
    model: model,
    options: {
      temperature: 0.5,
      repeat_last_n: 2
    },
    messages: messages,
    stream: false
}).then((data) => {
    console.log(data.message.content)
})


