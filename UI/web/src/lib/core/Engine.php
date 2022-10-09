<?php declare(strict_types=1);

namespace webui\lib\core;

abstract class Engine {
    protected Configuration $config;

    protected function __construct(Configuration $config) {
        $this->config = $config;
    }

    abstract static public function make(): Engine;


}
