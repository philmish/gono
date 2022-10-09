<?php declare(strict_types=1);

use PHPUnit\Framework\TestCase;
use webui\app\App;

/**
 * @covers webui\app\App  
 * */
final class AppTest extends TestCase {

    public function testMakeSuccess(): void {
        putenv("WEBUI_DSN=test");
        App::make();
    }
}

