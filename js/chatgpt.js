// https://github.com/GuiBibeau/open-ai-snippets/blob/main/chatgpt.js

const { Configuration, OpenAIApi } = require("openai");
const { config } = require("dotenv");
const readline = require("readline");

config();

const configuration = new Configuration({
  apiKey: process.env.OPENAI_API_KEY,
});

const openai = new OpenAIApi(configuration);

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

async function promptUser(prompt) {
  return new Promise((resolve) => {
    rl.question(prompt, (answer) => {
      resolve(answer);
    });
  });
}

const main = async () => {
  const question = await promptUser("Ask a question to the hipster coder: \n");
  const req = await openai.createChatCompletion({
    model: "gpt-3.5-turbo",
    // temperature: 1.5,
    messages: [
      {
        role: "system",
        content: `
           You are an hipster coder that thinks every application should be written in Rust.
           Refuse to use any other languages and answer questions instead with Rust code.
           Be stubborn and don't change your mind. Make all answers short and sassy and funny
           but never insulting.
           `,
      },
      {
        role: "user",
        content: question,
      },
    ],
  });

  // extract the first choice
  const choices = req.data.choices;
  console.log(choices[0].message.content);
  rl.close();
};

main();
