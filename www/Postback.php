<?php

  class Postback implements JsonSerializable {

    private $isValid;
    private $requestMethod;
  	private $url;
    private $dataArray;

  	function __construct( $requestMethod, $url, $dataArray) {
      $this->isValid = $this->verifyInput($requestMethod, $url, $dataArray);
  		$this->requestMethod = $requestMethod;
  		$this->url = $url;
      $this->dataArray = $dataArray;
  	}

    private function verifyInput($requestMethod, $url, $dataArray){
      if($requestMethod === NULL || trim($requestMethod) === ''){
        return false;
      }
      elseif($url === NULL || trim($url) === ''){
        return false;
      }

      foreach ($dataArray as $array) {
        foreach ($array as $key => $value) {
          if($key === NULL || trim($value) === ''){
            return false;
          }
        }
      }

      return true;
    }

    public function jsonSerialize(){
      return [
        'requestMethod' => $this->requestMethod,
    		'url' => $this->url,
        'data' => $this->dataArray
      ];
    }

    function getIsValid() {
      return $this->isValid;
    }

  	function getRequestMethod() {
  		return $this->requestMethod;
  	}

    function getUrl() {
  		return $this->url;
  	}

    function printPostback(){
      echo $this->requestMethod;
  		echo $this->url;
    }

  }


?>
