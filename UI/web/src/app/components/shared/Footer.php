<?php declare(strict_types=1);

namespace webui\app\components\shared;

use webui\lib\core\Component;

final class Footer implements Component {
    
    public function __construct(?string $js) {
        $this->js = $js;
    }

    public function getHTML(): string {
        $html = <<<EOT
            <script src={$this->js}></script>
        </body>
        </html>
        EOT;
        return $html;
    }
}
