<?php declare(strict_types=1);

namespace webui\app;

use webui\lib\core\Component;
use webui\lib\core\DataProvider;

abstract class ApiProvider implements DataProvider {

    private ?string $dsn;

    public function __construct() {
        $this->dsn = NULL;
    }

    public function setDSN(string $dsn): void {
        $this->dsn = $dsn;
    }
    
    abstract protected function provide(?array $args = []): array;
    abstract public function generateComponent(): Component;
}
