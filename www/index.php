<?php

require "Postback.php";

$uri = parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH);
$uri = explode( '/', $uri );

$redis = new Redis();

$redis->connect('redis', 6379);
$params = (array) json_decode(file_get_contents('php://input'), TRUE);

$endpoint = $params['endpoint']['url'];
$method = $params['endpoint']['method'];
$inputDataArray = $params['data'];

print_r($inputDataArray[0]);

$postback = new Postback($method, $endpoint,  $inputDataArray);

if($postback->getIsValid()){
  //push postback object to redis stack
  $redis->lpush("postback-list", json_encode($postback));
}

?>
