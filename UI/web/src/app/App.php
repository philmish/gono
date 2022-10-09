<?php declare(strict_types=1);

namespace webui\app;

use webui\lib\core\Engine;

final class App extends Engine {
    
    static public function make(): self {
        $config = new AppConfiguration(
            AppConfiguration::getFromEnv("WEBUI_CSS", "http://localhost:8081/static/style.css"),
            AppConfiguration::getFromEnv("WEBUI_JS", "http://localhost:8081/static/script.js")
        );
        $config->addViews();
        $app = new self($config);
        $app->config->setDSN(
            AppConfiguration::getFromEnv("WEBUI_DSN", "http://backend:7195/"), 
        );
        return $app;
    }

    public function serve(string $uri): void {
        switch ($uri) {
        case "/":
            $home = $this->config->getTemplate("home");
            if (!$home) {
                echo "Not found";
                return;
            }
            $home->render(NULL);
        }
    }

}


