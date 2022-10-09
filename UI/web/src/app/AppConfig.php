<?php declare(strict_types=1);

namespace webui\app;

use webui\app\views\HomeView;
use webui\lib\core\Configuration;

final class AppConfiguration extends Configuration {

    private string $css;
    private string $js;

    public function __construct(string $css, string $js) {
        $this->css = $css;
        $this->js = $js;
        parent::__construct();
    }

    public function addViews(): void {
        $sites = [
            "home" => new HomeView($this->css, $this->js),
        ];
        foreach ($sites as $site => $template) {
            $this->addTemplate($site, $template);
        }
    }

    public function injectDSN(ApiProvider $db): ApiProvider {
        $db->setDSN($this->dsn);
        return $db;
    }

}
