<?php declare(strict_types=1);

namespace webui\app\components\shared;

use webui\lib\core\Component;

final class Head implements Component {
    
    public function __construct(?string $css) {
        $this->css = $css;
    }

    public function getHTML(): string {
        $html = <<<EOT
        <!DOCTYPE html>
        <head>
            <meta charset="utf-8">
            <meta name="keywords" content="example, html, head, meta">
            <link rel="stylesheet" href={$this->css}>
            <title>Gono WebUI</title>
        </head>
        <body>
        EOT;
        return $html;
    }
}
