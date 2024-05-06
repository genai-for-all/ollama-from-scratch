//let ollamaUrl = "http://bob.local:11434"
let ollamaUrl = "http://host.docker.internal:11434"
//let model = "deepseek-coder"
let model = "phi3"

async function postData(url = "", data = {}) {
    var responseText = ""
    try {
        const response = await fetch(url, {
          method: "POST",
          cache: "no-cache",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        })
    
        const reader = response.body.getReader()
    
        while (true) {
          const { done, value } = await reader.read()
          if (done) {    
            responseText = responseText + "\n"
            return
          }
          // Otherwise do something here to process current chunk
          const decodedValue = new TextDecoder().decode(value)
          const part = JSON.parse(decodedValue)
          process.stdout.write(part.message.content)
          responseText = responseText + decodedValue
        }
    
      } catch(error) {
        console.log("😡", error)
      }
      return responseText
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
    stream: true
}).then((data) => {
    console.log("📝", data)
})
