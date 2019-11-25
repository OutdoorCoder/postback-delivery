<?php

require "Postback.php";
//code seems to be no longer needed now that apache is hosting the server
//header("Access-Control-Allow-Origin: *");
//header("Content-Type: application/json; charset=UTF-8");
//header("Access-Control-Allow-Methods: OPTIONS,GET,POST,PUT,DELETE");
//header("Access-Control-Max-Age: 3600");
//header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");

$uri = parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH);
$uri = explode( '/', $uri );


$requestMethod = $_SERVER["REQUEST_METHOD"];
$redis = new Redis();

switch ($requestMethod) {
    case 'GET':
        echo 'GET';

        //I don't think I'll recieve GET requests in this app.
        //Instructions just mention post requests, which makes sense.
        //TODO: remove this before the end of the project
        try {
            echo 'here';
        } catch (\Exception $e) {
            var_dump($e->getMessage())  ;
            die;
        }

        break;
    case 'POST':

        $redis->connect('redis', 6379);
        $params = (array) json_decode(file_get_contents('php://input'), TRUE);

        $endpoint = $params['endpoint']['url'];
        $method = $params['endpoint']['method'];
        $inputDataArray = $params['data'];

        for($x = 0; $x < count($params['data']); $x++){

          $postback = new Postback($method, $endpoint,  $inputDataArray[$x]['mascot'], $inputDataArray[$x]['location']);

          if($postback->getIsValid()){
            //push postback object to redis stack
            $redis->lpush("postback-list", serialize($postback));
          }
        }

        $arList = $redis->lrange("postback-list", 0 ,6);
        print_r($arList);

        break;
      default:
        //Just ignore the request if this happens, maybe log it.
        echo 'bad';
      }

// pass the request method and user ID to the PersonController and process the HTTP request:

?>
