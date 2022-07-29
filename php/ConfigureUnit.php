<?php	

// https://gist.github.com/peterfox/bb9c5b4e084fd92e3ba7d253fe726b1f
// A console command for applying a config to Nginx Unit 

namespace App\Console\Commands\Nginx;

use Illuminate\Console\Command;

class ConfigureUnit extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'nginx:unit:configure {--control=/usr/local/var/run/unit/control.sock : The path to the control socket}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Command description';

    /**
     * Execute the console command.
     */
    public function handle()
    {
        $basePath = base_path();
        $controlSocket = $this->option('control');
        $json = <<<JSON
{
    "listeners": {
        "*:80": {
            "pass": "routes"
        }
    },

    "routes": [
        {
            "match": {
                "uri": "!/index.php"
            },
            "action": {
                "share": "$basePath/public\$uri",
                "fallback": {
                    "pass": "applications/laravel"
                }
            }
        }
    ],

    "applications": {
        "laravel": {
            "type": "php",
            "root": "$basePath/public/",
            "script": "index.php",
            "processes": {}
        }
    }
}
JSON;

        $data =\Http::withOptions([
            'curl' => [
                CURLOPT_UNIX_SOCKET_PATH => $controlSocket,
            ],
        ])
            ->put('http://localhost/config/', json_decode($json, flags: JSON_THROW_ON_ERROR))
            ->json();

        if (! isset($data['success'])) {
            $this->components->error($data['error'] ?? 'Unknown error');
            return self::FAILURE;
        }

        $this->components->info('Configured!');
    }
}
