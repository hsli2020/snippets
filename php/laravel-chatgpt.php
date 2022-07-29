<?php  https://github.com/beyondcode/chatgpt-clone/blob/main/routes/web.php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use OpenAI\Laravel\Facades\OpenAI;

Route::get('/', function () {
    $messages = collect(session('messages', []))
        ->reject(fn ($message) => $message['role'] === 'system');
    return view('welcome', [ 'messages' => $messages ]);
});

Route::post('/', function (Request $request) {
    $messages = $request->session()->get('messages', [
        ['role' => 'system', 
         'content' => 'You are LaravelGPT - A ChatGPT clone. Answer as concisely as possible.']
    ]);

    $messages[] = ['role' => 'user', 'content' => $request->input('message')];

    $response = OpenAI::chat()->create([
        'model' => 'gpt-3.5-turbo',
        'messages' => $messages
    ]);

    $messages[] = ['role' => 'assistant', 'content' => $response->choices[0]->message->content];

    $request->session()->put('messages', $messages);
    return redirect('/');
});

Route::get('/reset', function (Request $request) {
    $request->session()->forget('messages');
    return redirect('/');
});

<body class="antialiased">
  <div class="flex flex-col space-y-4 p-4">
  @foreach($messages as $message)
    <div class="flex rounded-lg p-4 
      @if ($message['role'] === 'assistant') bg-green-200 flex-reverse @else bg-blue-200 @endif">
      <div class="ml-4">
        <div class="text-lg">
          @if ($message['role'] === 'assistant')
            <a href="#" class="font-medium text-gray-900">LaravelGPT</a>
          @else
            <a href="#" class="font-medium text-gray-900">You</a>
          @endif
        </div>
        <div class="mt-1">
          <p class="text-gray-600">
            {!! \Illuminate\Mail\Markdown::parse($message['content']) !!}
          </p>
        </div>
      </div>
    </div>
  @endforeach
  </div>

  <form class="p-4 flex space-x-4 justify-center items-center" action="/" method="post">
      @csrf
      <label for="message">Laravel Question:</label>
      <input id="message" type="text" name="message" class="border rounded-md  p-2 flex-1" />
      <a class="bg-gray-800 text-white p-2 rounded-md" href="/reset">Reset Conversation</a>
  </form>
</body>
