<?php

namespace App\Http\Controllers;

use App\Models\Transcript;
use Illuminate\Http\Request;

class DownloadTranscriptController extends Controller
{
    public function __invoke(Transcript $transcript, Request $request)
    {
        $this->authorize('view', $transcript);

        $transcriptVtt = $transcript->translations[$request->get('language')] ?? $transcript->transcript;

        return response($transcriptVtt, 200, [
            'Content-Type' => 'text/vtt',
            'Content-Disposition' => 'attachment; filename="transcript.vtt"',
            'Content-Length' => strlen($transcriptVtt),
        ]);
    }
}

<?php

namespace App\Http\Controllers;

use App\Models\User;
use Laravel\Socialite\Facades\Socialite;

class GitHubLoginController extends Controller
{
    public function redirect()
    {
        return Socialite::driver('github')->redirect();
    }

    public function callback()
    {
        $user = Socialite::driver('github')->user();

        $user = User::firstOrCreate([
            'github_id' => $user->getId(),
        ], [
            'name' => $user->getName(),
            'email' => $user->getEmail(),
            'github_username' => $user->getNickname(),
        ]);

        auth()->login($user);

        return redirect()->action(NewTranscriptionController::class);
    }
}

<?php

namespace App\Http\Controllers;

class LogoutController extends Controller
{
    public function __invoke()
    {
        auth()->logout();
        return redirect()->to('/');
    }
}
