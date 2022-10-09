<?php declare(strict_types=1);

namespace webui\lib\core;

interface DataProvider {
    public function generateComponent(): Component;
}


