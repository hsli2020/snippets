###########################################################
# https://github.com/ddddddeon/a
# CLI tool to generate code from GPT3 
# Usage
#
# You will need an OpenAI API key, and to set the environment variable OPENAI_API_KEY.
#
# Invoke the a command followed by a prompt. If the first word in the prompt is a 
# programming language or file-format the pretty-printer recognizes, it will syntax 
# highlight the output.
#
# $ a python script that fetches a url
# $ a rust program that showcases its various features
# $ a yaml manifest describing a kubernetes deployment
###########################################################

# main.rs

pub mod gpt;
pub mod util;

fn main() {
    let mut args: Vec<_> = std::env::args().collect();
    args.remove(0);

    if args.is_empty() {
        println!("no prompt provided");
        std::process::exit(1);
    }

    let lang = args[0].clone();
    let prompt = args.join(" ");
    let api_key = std::env::var("OPENAI_API_KEY")
        .expect("Please set the OPENAI_API_KEY environment variable");

    let client = gpt::GPTClient::new(api_key);
    let mut response = client
        .prompt(prompt)
        .expect("Could not make request to API");

    response.push('\n');
    if let Some(r) = response.strip_prefix("\n\n") {
        response = String::from(r);
    }

    #[cfg(feature = "clipboard")]
    {
        util::copy_to_clipboard(&response);
    }

    util::pretty_print(&response, &lang);
}

# gpt.rs

use super::util;
use reqwest::{
    blocking::Client,
    header::{HeaderMap, HeaderValue},
};
use serde::{Deserialize, Serialize};
use serde_json::{from_str, Value};
use std::error::Error;
use std::io::Read;

const OPENAI_API_URL: &str = "https://api.openai.com/v1/completions";
const OPENAI_MODEL: &str = "text-davinci-003";
const MAX_TOKENS: u32 = 4097;
const TEMPERATURE: f32 = 0.2;

type BoxResult<T> = Result<T, Box<dyn Error>>;

#[derive(Serialize, Deserialize, Debug)]
struct Prompt {
    model: String,
    prompt: String,
    temperature: f32,
    max_tokens: u32,
}

pub struct GPTClient {
    api_key: String,
    url: String,
}

impl GPTClient {
    pub fn new(api_key: String) -> Self {
        GPTClient {
            api_key,
            url: String::from(OPENAI_API_URL),
        }
    }

    pub fn prompt(&self, prompt: String) -> BoxResult<String> {
        let prompt_length = prompt.len() as u32;
        if prompt_length >= MAX_TOKENS {
            return Err(format!(
                "Prompt cannot exceed length of {} characters",
                MAX_TOKENS - 1
            )
            .into());
        }

        let p = Prompt {
            max_tokens: MAX_TOKENS - prompt_length,
            model: String::from(OPENAI_MODEL),
            prompt,
            temperature: TEMPERATURE,
        };

        let mut auth = String::from("Bearer ");
        auth.push_str(&self.api_key);

        let mut headers = HeaderMap::new();
        headers.insert("Authorization", HeaderValue::from_str(auth.as_str())?);
        headers.insert("Content-Type", HeaderValue::from_str("application/json")?);

        let body = serde_json::to_string(&p)?;

        let client = Client::new();
        let mut res = client.post(&self.url).body(body).headers(headers).send()?;

        let mut response_body = String::new();
        res.read_to_string(&mut response_body)?;
        let json_object: Value = from_str(&response_body)?;
        let answer = json_object["choices"][0]["text"].as_str();

        match answer {
            Some(a) => Ok(String::from(a)),
            None => {
                util::pretty_print(&response_body, "json");
                Err(format!("JSON parse error").into())
            }
        }
    }
}

# src/util.rs

use bat::PrettyPrinter;
use bat::Syntax;
use copypasta_ext::prelude::*;
use copypasta_ext::x11_fork::ClipboardContext;

fn lang_exists(lang: &str, langs: &Vec<Syntax>) -> bool {
    for l in langs {
        if l.name.to_lowercase() == lang.to_lowercase() {
            return true;
        }
        for e in &l.file_extensions {
            if e == &lang.to_lowercase() {
                return true;
            }
        }
    }
    false
}

pub fn pretty_print(str: &str, lang: &str) {
    let mut lang = lang.to_owned();
    let mut pp = PrettyPrinter::new();

    let langs: Vec<_> = pp.syntaxes().collect();
    if !lang_exists(&lang, &langs) {
        lang = "txt".to_owned();
    }

    pp.input_from_bytes(str.as_bytes())
        .language(&lang)
        .print()
        .unwrap();
}

pub fn copy_to_clipboard(str: &str) {
    let mut ctx = ClipboardContext::new().unwrap();
    ctx.set_contents(str.to_owned()).unwrap();
}
