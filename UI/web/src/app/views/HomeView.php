<?php declare(strict_types=1);

namespace webui\app\views;

use webui\app\templates\BaseTemplate;
use webui\lib\core\DataProvider;

final class HomeView extends BaseTemplate {

    protected function composeBody(?DataProvider $db, array $args = []): string {
        return <<<EOT
            <h1>Hello World</h1>
        EOT;
    }
    
}


