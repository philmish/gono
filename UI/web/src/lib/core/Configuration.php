<?php declare(strict_types=1);

namespace webui\lib\core;

abstract class Configuration {
    private array $siteMap;
    private ?string $dsn;

    public function __construct(array $siteMap = [], ?string $dsn = NULL) {
        $this->siteMap = $siteMap;
        $this->dsn = $dsn;
    }

    public function setDSN(string $dsn) {
        $this->dsn = $dsn;
    }

    public function addTemplate(string $key, Template $template): void {
        $this->siteMap[$key] = $template;
    }

    public function getTemplate(string $key): ?Template {
        if (!array_key_exists($key, $this->siteMap)) {
            return NULL;
        }
        return $this->siteMap[$key];
    }

    static public function getFromEnv(string $key, string $default): string {
        $var = getenv($key);
        if (!$var) {
            return $default;
        }
        return $var;
    }
}
