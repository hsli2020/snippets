// https://ahmadrosid.com/blog/laravel-openai-streaming-response
// https://github.com/ahmadrosid/laravel-openai-streaming

// Listening Server-Sent Event with Javascript

<script>
  const form = document.querySelector("form");
  const result = document.getElementById("result");

  form.addEventListener("submit", (event) => {
    event.preventDefault();
    const input = event.target.input.value;
    if (input === "") return;
    const question = document.getElementById("question");
    question.innerText = input;
    event.target.input.value = "";

    const queryQuestion = encodeURIComponent(input);
    const source = new EventSource("/ask?question=" + queryQuestion);
    source.addEventListener("update", function (event) {
      if (event.data === "<END_STREAMING_SSE>") {
        source.close();
        return;
      }
      result.innerText += event.data;
    });
  });
</script>

<?php

// routes/web.php:

use App\Http\Controllers\AskController;
use Illuminate\Support\Facades\Route;

Route::get('/', function () { return view('welcome'); });
Route::get("/ask", AskController::class);

// app/Http/Controllers/AskController.php:

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use OpenAI\Laravel\Facades\OpenAI;

class AskController extends Controller
{
    public function __invoke(Request $request)
    {
        $question = $request->query('question');
        return response()->stream(function () use ($question) {
            $stream = OpenAI::completions()->createStreamed([
                'model' => 'text-davinci-003',
                'prompt' => $question,
                'max_tokens' => 1024,
            ]);

            foreach ($stream as $response) {
                $text = $response->choices[0]->text;
                if (connection_aborted()) {
                    break;
                }

                echo "event: update\n";
                echo 'data: ' . $text;
                echo "\n\n";
                ob_flush();
                flush();
            }

            echo "event: update\n";
            echo 'data: <END_STREAMING_SSE>';
            echo "\n\n";
            ob_flush();
            flush();
        }, 200, [
            'Cache-Control' => 'no-cache',
            'X-Accel-Buffering' => 'no',
            'Content-Type' => 'text/event-stream',
        ]);
    }
}

// Nginx Config

// When deploying your service with Nginx, it is crucial to configure your Nginx
// configuration file correctly. Specifically, you should unset the Connection
// header and set proxy_http_version to 1.1.

location ^~ /ask$ {
    proxy_http_version 1.1;
    add_header Connection '';

    fastcgi_pass unix:/var/run/php/php8.1-fpm.sock;
    fastcgi_index index.php;
    fastcgi_param SCRIPT_FILENAME $realpath_root$fastcgi_script_name;
    include fastcgi_params;
}
