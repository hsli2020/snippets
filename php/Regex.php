<?php // https://github.com/chevere/regex/blob/1.0/src/Regex.php

/*
 * This file is part of Chevere.
 *
 * (c) Rodolfo Berrios <rodolfo@chevere.org>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

declare(strict_types=1);

namespace Chevere\Regex;

use Chevere\Regex\Exceptions\NoMatchException;
use Chevere\Regex\Interfaces\RegexInterface;
use InvalidArgumentException;
use LogicException;
use function Chevere\Message\message;

final class Regex implements RegexInterface
{
    private string $noDelimiters;

    private string $noDelimitersNoAnchors;

    /**
     * @throws InvalidArgumentException
     */
    public function __construct(
        private string $pattern
    ) {
        $this->assertPattern();
        $delimiter = $this->pattern[0];
        $this->noDelimiters = trim($this->pattern, $delimiter);
        $this->noDelimitersNoAnchors = strval(
            preg_replace('#^\^(.*)\$$#', '$1', $this->noDelimiters)
        );
    }

    public function __toString(): string
    {
        return $this->pattern;
    }

    public function noDelimiters(): string
    {
        return $this->noDelimiters;
    }

    public function noDelimitersNoAnchors(): string
    {
        return $this->noDelimitersNoAnchors;
    }

    public function match(string $value): array
    {
        $match = @preg_match($this->pattern, $value, $matches);
        if (is_int($match)) {
            return $match === 1 ? $matches : [];
        }
        // @codeCoverageIgnoreStart
        throw new LogicException(
            (string) message(
                'Error `%function%` %error%',
                function: 'preg_match',
                error: static::ERRORS[preg_last_error()],
            )
        );
        // @codeCoverageIgnoreEnd
    }

    public function assertMatch(string $value): void
    {
        if ($this->match($value)) {
            return;
        }

        throw new NoMatchException(
            (string) message(
                'String `%string%` does not match regex `%pattern%`',
                pattern: $this->pattern,
                string: $value,
            ),
            100
        );
    }

    public function matchAll(string $value): array
    {
        $match = @preg_match_all($this->pattern, $value, $matches);
        if (is_int($match)) {
            return $match === 1 ? $matches : [];
        }
        // @codeCoverageIgnoreStart
        throw new LogicException(
            (string) message(
                'Error `%function%` %error%',
                function: 'preg_match',
                error: static::ERRORS[preg_last_error()],
            )
        );
        // @codeCoverageIgnoreEnd
    }

    public function assertMatchAll(string $value): void
    {
        if ($this->matchAll($value)) {
            return;
        }

        throw new NoMatchException(
            (string) message(
                'String `%string%` does not match all `%pattern%`',
                pattern: $this->pattern,
                string: $value,
            ),
            110
        );
    }

    private function assertPattern(): void
    {
        if (@preg_match($this->pattern, '') !== false) {
            return;
        }

        throw new InvalidArgumentException(
            (string) message(
                'Invalid regex pattern `%pattern%` provided: %error%',
                pattern: $this->pattern,
                error: static::ERRORS[preg_last_error()],
            )
        );
    }
}
