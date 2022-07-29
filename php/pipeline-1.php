<?php

interface PipelineStage {
    public function process(mixed $payload): mixed;
}

class Pipeline {
    private array $stages = [];

    public function addStage(PipelineStage $stage): self {
        $this->stages[] = $stage;
        return $this;
    }

    public function process(mixed $payload): mixed {
        return array_reduce(
            $this->stages,
            fn($carry, $stage) => $stage->process($carry),
            $payload
        );
    }
}

class TrimWhitespaceStage implements PipelineStage {
    public function process(mixed $payload): mixed {
        return is_string($payload) ? trim($payload) : $payload;
    }
}

class ConvertToUppercaseStage implements PipelineStage {
    public function process(mixed $payload): mixed {
        return is_string($payload) ? strtoupper($payload) : $payload;
    }
}

class AddPrefixStage implements PipelineStage {
    private string $prefix;

    public function __construct(string $prefix) {
        $this->prefix = $prefix;
    }

    public function process(mixed $payload): mixed {
        return is_string($payload) ? $this->prefix . $payload : $payload;
    }
}

$pipeline = (new Pipeline())
    ->addStage(new TrimWhitespaceStage())
    ->addStage(new ConvertToUppercaseStage())
    ->addStage(new AddPrefixStage("Hello, "));

$input = "  world ";
$output = $pipeline->process($input);

echo $output; // Output: "HELLO, WORLD"
