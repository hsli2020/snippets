#!/usr/bin/env python3

# https://github.com/reorx/ai.py

import os
import re
import json
import logging
import argparse

from typing import Optional, Tuple, Union, Callable
from urllib import request, parse
from urllib.error import HTTPError
from http.client import HTTPResponse, IncompleteRead

__version__ = '0.1.0'

class Config:
    api_key = os.environ.get('AI_CLI_API_KEY')
    default_params = {
        # 'max_tokens': 80,
        # 'temperature': 0.8,
        # 'top_p': 1,
        # 'frequency_penalty': 0.5,
        # 'presence_penalty': 0,
    }
    default_model = 'gpt-3.5-turbo'
    timeout = None
    verbose = False
    debug = False


lg = logging.getLogger(__name__)

home = os.path.expanduser('~')


def main():
    # the `formatter_class` can make description & epilog show multiline
    parser = argparse.ArgumentParser(description="A simple CLI for ChatGPT API", epilog="", formatter_class=argparse.RawDescriptionHelpFormatter)

    # arguments
    parser.add_argument('prompt', metavar="PROMPT", type=str, nargs='?', help="your prompt, leave it empty to run REPL")

    # options
    parser.add_argument('-s', '--system', type=str, help="system message to use at the beginning of the conversation. if starts with @, the message will be located through ~/.ai_cli_prompts.json")
    parser.add_argument('-v', '--verbose', action='store_true', help="verbose mode, show params and role name")
    parser.add_argument('-d', '--debug', action='store_true', help="debug mode, enable logging")
    # parser.add_argument('-i', '--stdin', action='store_true', help="read prompt from stdin")

    # --version
    parser.add_argument('--version', action='version',
        version='%(prog)s {version}'.format(version=__version__))

    args = parser.parse_args()

    # config
    Config.verbose = args.verbose
    Config.debug = args.debug
    if Config.debug:
        logging.basicConfig(level=logging.DEBUG)
    if not Config.api_key:
        # try to load from file ~/.ai_cli_api_key
        api_key_file = os.path.join(home, '.ai_cli_api_key')
        if os.path.exists(api_key_file):
            with open(api_key_file) as f:
                Config.api_key = f.read().strip()
        else:
            print(red('ERROR: missing API key'))
            print('Please set the environment variable AI_CLI_API_KEY or create a file ~/.ai_cli_api_key')
            exit(1)

    # load prompts
    prompts_store = PromptsStore()
    prompts_store.load_from_file()

    # create session
    system_prompt = args.system
    if system_prompt.startswith('@'):
        system_prompt = prompts_store.get('system', system_prompt[1:])

    session = ChatSession(Config.api_key, messages=new_messages(system_prompt))
    if args.verbose:
        for i in session.messages:
            print_message(i)

    # call the function
    if args.prompt:
        chat_once(session, args.prompt)
    else:
        repl(session)


def chat_once(session, prompt):
    try:
        res_message = session.chat(prompt)
    except TimeoutError:
        print(red('ERROR: timeout'))
        return
    print_message(res_message)


def repl(session):
    pass


inline_code_re = re.compile(r'`([^\n`]+)`')
multiline_code_re = re.compile(r'```\w*\n([^`]+)\n```')


def print_message(message):
    role = message['role']
    role_with_padding = f' {role} '
    content = message['content']

    # find inline code and replace with color
    content = multiline_code_re.sub(lambda m: m.group(0).replace(m.group(1), cyan(m.group(1))), content)
    content = inline_code_re.sub(lambda m: m.group(0).replace(m.group(1), cyan(m.group(1))), content)

    content_color = lambda s: s
    role_color = white_hl
    if role == 'system':
        content_color = yellow
        role_color = yellow_hl
    elif role == 'user':
        content_color = green
        role_color = green_hl

    s = content_color(content)
    if (Config.verbose):
        s = f'{role_color(role_with_padding)} {s}'

    print(s + '\n')


def new_messages(system_message):
    if system_message:
        return [{
            'role': 'system',
            'content': system_message,
        }]
    return []


# Prompts #

class PromptsStore:
    def __init__(self):
        self.data = {}

    def load_from_file(self):
        prompts_file = os.path.join(home, '.ai_cli_prompts.json')
        if os.path.exists(prompts_file):
            with open(prompts_file) as f:
                self.data = json.load(f)

    def get(self, role, name):
        return self.data.get(role, {})[name]


# Session #

OPENAI_BASE_URL = 'https://api.openai.com/v1/'

class ChatSession:
    def __init__(self, api_key, messages=None):
        self.api_key = api_key
        if messages is None:
            messages = []
        self.messages = messages

    def chat(self, content, params=None):
        message = {
            'role': 'user',
            'content': content,
        }
        self.messages.append(message)
        return self.create_completion(params=params)

    def create_completion(self, params=None) -> dict:
        url = f'{OPENAI_BASE_URL}chat/completions'
        headers = {
            'Authorization': f'Bearer {self.api_key}',
        }

        if not params:
            params = dict(Config.default_params)
        data = dict(params)
        data.update(
            model=Config.default_model,
            messages=self.messages,
        )

        try:
            res, body_b = http_request('POST', url, headers=headers, data=data, logger=lg, timeout=Config.timeout)
        except HTTPError as e:
            raise RequestError(e.status, e.read().decode()) from None
        res_data = json.loads(body_b)
        res_message = res_data['choices'][0]['message']

        self.messages.append(res_message)
        return res_message


# HTTP request #

def http_request(method, url, params=None, headers=None, data: Optional[Union[dict, list, bytes]] = None, timeout=None, logger=None) -> Tuple[HTTPResponse, bytes]:
    if params:
        url = f'{url}?{parse.urlencode(params)}'
    if not headers:
        headers = {}
    if data and isinstance(data, (dict, list)):
        data = json.dumps(data, ensure_ascii=False).encode()
        if 'Content-Type' not in headers:
            headers['Content-Type'] = 'application/json; charset=utf-8'
    if logger:
        logger.debug(f'request: {method} {url}\nheaders: {headers}\ndata: {data}')
    req = request.Request(url, method=method, headers=headers, data=data)
    res = request.urlopen(req, timeout=timeout)  # raises: (HTTPException, urllib.error.HTTPError)
    try:
        body_b: bytes = res.read()
    except IncompleteRead as e:
        body_b: bytes = e.partial
    if logger:
        logger.debug(f'response: {res.status}, {body_b}')
    return res, body_b


class RequestError(Exception):
    def __init__(self, status, body) -> None:
        self.status = status
        self.body = body

    def __str__(self):
        return f'{self.__class__.__name__}: {self.status}, {self.body}'


# Color #

def esc(*codes: Union[int, str]) -> str:
    """Produces an ANSI escape code from a list of integers
    :rtype: text_type
    """
    return '\x1b[{}m'.format(';'.join(str(c) for c in codes))


def make_color(start, end: str) -> Callable[[str], str]:
    def color_func(s: str) -> str:
        return start + s + end
    return color_func


END = esc(0)

FG_END = esc(39)
black = make_color(esc(30), FG_END)
red = make_color(esc(31), FG_END)
green = make_color(esc(32), FG_END)
yellow = make_color(esc(33), FG_END)
blue = make_color(esc(34), FG_END)
magenta = make_color(esc(35), FG_END)
cyan = make_color(esc(36), FG_END)
white = make_color(esc(37), FG_END)

BG_END = esc(49)
black_bg = make_color(esc(40), BG_END)
red_bg = make_color(esc(41), BG_END)
green_bg = make_color(esc(42), BG_END)
yellow_bg = make_color(esc(43), BG_END)
blue_bg = make_color(esc(44), BG_END)
magenta_bg = make_color(esc(45), BG_END)
cyan_bg = make_color(esc(46), BG_END)
white_bg = make_color(esc(47), BG_END)

HL_END = esc(22, 27, 39)
#HL_END = esc(22, 27, 0)

black_hl = make_color(esc(1, 30, 7), HL_END)
red_hl = make_color(esc(1, 31, 7), HL_END)
green_hl = make_color(esc(1, 32, 7), HL_END)
yellow_hl = make_color(esc(1, 33, 7), HL_END)
blue_hl = make_color(esc(1, 34, 7), HL_END)
magenta_hl = make_color(esc(1, 35, 7), HL_END)
cyan_hl = make_color(esc(1, 36, 7), HL_END)
white_hl = make_color(esc(1, 37, 7), HL_END)

bold = make_color(esc(1), esc(22))
italic = make_color(esc(3), esc(23))
underline = make_color(esc(4), esc(24))
strike = make_color(esc(9), esc(29))
blink = make_color(esc(5), esc(25))

if __name__ == '__main__':
    main()
