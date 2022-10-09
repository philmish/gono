<?php declare(strict_types=1);

namespace webui;
require_once('src/autoload.php');

use webui\app\App;

$app = App::make();
$app->serve("/");


