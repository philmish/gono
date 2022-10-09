<?php declare(strict_types=1);

namespace webui\app\components\shared;

use webui\lib\core\Component;

final class Navbar implements Component  {
    
    public function getHTML(): string {
        $html = <<<EOT
            <div class="navbar">
                <a class="navitem" href="#">Projects</a>
            </div>
        EOT;
        return $html;
    }
}
