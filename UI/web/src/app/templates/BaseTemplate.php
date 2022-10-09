<?php declare(strict_types=1);

namespace webui\app\templates;

use webui\app\components\shared\Footer;
use webui\app\components\shared\Head;
use webui\lib\core\DataProvider;
use webui\lib\core\Template;

abstract class BaseTemplate implements Template {

    private string $css;
    private string $js;

    public function __construct(string $css, string $js) {
        $this->css = $css;
        $this->js = $js; 
    }

    private function head(): string {
        $h = new Head($this->css);
        return $h->getHTML();
    }

    private function foot(): string {
        $f = new Footer($this->js);
        return $f->getHTML();
    }

    abstract protected function composeBody(?DataProvider $db, array $args = []): string;

    public function render(?DataProvider $db, array $args = []): void {
        $body = $this->composeBody($db, $args);
        echo <<<EOT
        {$this->head()}
            {$body}
        {$this->foot()}
        EOT;
    }
}
