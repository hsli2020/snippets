# composer require openai-php/client

# generate-test.php

<?php

require __DIR__ . '/vendor/autoload.php';

$client = OpenAI::client('<YOUR_API_KEY>');

$result = $client->completions()->create([
    'model' => 'text-davinci-003',
#   'model' => 'code-davinci-002',
    'prompt' => file_get_contents(__DIR__ . '/prompt.txt'),
    'max_tokens' => 1000
]);

echo $result['choices'][0]['text'];
?>

# prompt.txt

Generate PHPUnit test for a "someMagic()" method with data provider of 4 use cases for this code:

'''php
<?php

class SomeClass
{
    public function someMagic(int $firstNumber, int $secondNumber)
    {
        if ($firstNumber > 10) {
            return $firstNumber * $secondNumber;
        }

        return $firstNumber + $secondNumber;
    }
}
'''

php generate-test.php

<?php

use PHPUnit\Framework\TestCase;

class SomeClassTest extends TestCase
{
    /**
     * @covers SomeClass::someMagic
     * @dataProvider someDataProvider
     */
    public function testSomeMagic($firstNumber, $secondNumber, $expected)
    {
        $someClass = new SomeClass();
        $result = $someClass->someMagic($firstNumber, $secondNumber);
        $this->assertEquals($expected, $result);
    }

    public function someDataProvider()
    {
        return [
            [5, 5, 10],
            [12, 5, 60],
            [15, 10, 150],
            [20, 30, 600],
        ];
    }
}
