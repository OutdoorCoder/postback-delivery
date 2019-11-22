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


// the user id is, of course, optional and must be a number:

$requestMethod = $_SERVER["REQUEST_METHOD"];
$redis = new Redis();

switch ($requestMethod) {
    case 'GET':
        echo 'GET';


        try {
            $redis->connect('redis', 6379);
            echo 'here';
            echo $_SERVER['REQUEST_URI'];
        } catch (\Exception $e) {

            var_dump($e->getMessage())  ;
            die;
        }

        break;
    case 'POST':
        $params = (array) json_decode(file_get_contents('php://input'), TRUE);

        $endpoint = $params['endpoint']['url'];
        $method = $params['endpoint']['method'];
        $inputDataArray = $params['data'];
        $postback;

        for($x = 0; $x < count($params['data']); $x++){

          $postback = new Postback($method, $endpoint,  $inputDataArray[$x]['mascot'], $inputDataArray[$x]['location']);
          //$postback->printPostback();
          echo (int)$postback->getIsValid();
        }
        //echo 'POST';
        //var_dump($_POST);
        break;
      default:
        echo 'bad';
      }

// pass the request method and user ID to the PersonController and process the HTTP request:

?>
