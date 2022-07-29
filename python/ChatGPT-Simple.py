# https://github.com/logankilpatrick/ChatGPT-Simple
# Build a simple locally hosted version of ChatGPT in less than 100 lines of code

import os

import openai
from dotenv import load_dotenv
from flask import Flask, render_template, request

load_dotenv()  # load env vars from .env file
openai.api_key = os.getenv("OPENAI_API_KEY")

app = Flask(__name__)


@app.route("/")
def index():
    return render_template("index.html")


@app.route("/get_response")
def get_response():
    message = request.args.get("message")
    completion = openai.ChatCompletion.create(
        # You can switch this to `gpt-4` if you have access to that model.
        model="gpt-3.5-turbo",
        messages=[{"role": "user", "content": message}],
    )
    response = completion["choices"][0]["message"]["content"]
    return response


if __name__ == "__main__":
    app.run(debug=True)


<!DOCTYPE html>
<html>
  <h1>Chat with Chan</h1>
  <div class="chatbox" id="chatbox"></div>
  <input type="text" id="message" placeholder="Type your message here...">
  <button id="send">Send</button>
  <script>
    var chatbox = document.getElementById("chatbox");
    var message = document.getElementById("message");
    var send = document.getElementById("send");

    send.addEventListener("click", function() {
      var userMessage = message.value;
      var userDiv = document.createElement("div");
      userDiv.className = "message user";
      userDiv.innerHTML = "<strong>You:</strong> " + userMessage;
      chatbox.appendChild(userDiv);

      message.value = "";

      fetch("/get_response?message=" + encodeURIComponent(userMessage))
        .then(function(response) { return response.text(); })
        .then(function(botMessage) {
          var botDiv = document.createElement("div");
          botDiv.className = "message bot";
          botDiv.innerHTML = "<strong>Chan:</strong> " + botMessage;
          chatbox.appendChild(botDiv);

          chatbox.scrollTop = chatbox.scrollHeight;
      });
    });
  </script>
</html>
