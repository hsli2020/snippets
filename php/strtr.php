<?php

// https://rodolfoberrios.com/2024/01/08/message/
// https://github.com/chevere/message/blob/1.0/src/Message.php

final class Message
{
    /**
     * @var array<string, string>
     */
    private array $replacements = [];

    public function __construct(
        private string $template,
        float|int|string|Stringable ...$translate
    ) {
        $array = [];
        foreach ($translate as $key => $value) {
            $value = strval($value);
            $array["%{$key}%"] = $value;
            $array["{{{$key}}}"] = $value;
            $array["{{ {$key} }}"] = $value;
        }
        $this->replacements = $array;
    }

    public function __toString(): string
    {
        return strtr($this->template, $this->replacements);
    }

    public function template(): string
    {
        return $this->template;
    }

    public function replacements(): array
    {
        return $this->replacements;
    }
}

function message(string $template, float|int|string|Stringable ...$translate)
{
    return new Message($template, ...$translate);
}

$template = "I'm the {{what}}, {{what}}, {{name}}";

echo message($template,
    what: 'miggida',
    name: 'Mac Daddy',
);

// I'm the miggida, miggida, Mac Daddy
