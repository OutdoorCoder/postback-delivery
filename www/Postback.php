<?php

  class Postback implements JsonSerializable {

    private $isValid;
    private $requestMethod;
  	private $url;
    private $mascot;
    private $location;

  	function __construct( $requestMethod, $url, $mascot, $location ) {
      $this->isValid = $this->verifyInput($requestMethod, $url, $mascot, $location);

  		$this->requestMethod = $requestMethod;
  		$this->url = $url;
      $this->mascot = $mascot;
      $this->location = $location;
  	}

    private function verifyInput($requestMethod, $url, $mascot, $location){
      if($requestMethod === NULL || trim($requestMethod) === ''){
        return false;
      }
      elseif($url === NULL || trim($url) === ''){
        return false;
      }
      elseif($mascot === NULL || trim($mascot) === ''){
        return false;
      }
      elseif($location === NULL || trim($location) === ''){
        return false;
      }

      return true;
    }

    public function jsonSerialize(){
      return [
        'requestMethod' => $this->requestMethod,
    		'url' => $this->url,
        'mascot' => $this->mascot,
        'location' => $this->location
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

    function getMascot() {
  		return $this->mascot;
  	}

    function getLocation() {
  		return $this->location;
  	}

    function printPostback(){
      echo $this->requestMethod;
  		echo $this->url;
      echo $this->mascot;
      echo $this->location;
    }

  }


?>
