<?php declare(strict_types=1);

namespace webui\lib\core;

interface Template {
    public function render(?DataProvider $db, array $args = []): void;
}
