<?php 
session_start();
if(isset($_POST['submit'])){
    if($_POST['captcha']!=$_SESSION['code']){
        echo "<center><span style=color:red;font-size:15px>Invalid Captcha Code</center></span>";
    }else{
        echo "<pre>";
        print_r($_POST);exit;
    }
}
?>
<html lang="en">
<head>
  <title>Contact Form With Captcha</title>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">  
  <!-- script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script -->
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
  <style type="text/css">
    #multistep_form fieldset:not(:first-of-type) { display: none; }
  </style>
</head>
<body>

<h3 class="text-success" align="center">Contact Form With Captcha</h3><br>
<div class="container">
  <div class="panel-group">
    <div class="panel panel-primary">
     <div class="panel-heading">Contact Form With Captcha</div>
        <form class="form-horizontal" method="post">         
          <div class="panel-body">                 
            <div class="form-group">
              <label class="control-label col-sm-2" for="Name">Name:</label>
              <div class="col-sm-5">
                <input type="text" class="form-control" id="name" name="name" required>
              </div>
            </div>
            <div class="form-group">
              <label class="control-label col-sm-2" for="email">Email:</label>
              <div class="col-sm-5"> 
                <input type="email" class="form-control" id="email" name="email" required>
              </div>
            </div>  
            <div class="form-group">
              <label class="control-label col-sm-2" for="mobno">Mobile Number:</label>
              <div class="col-sm-5">
                <input type="number" class="form-control" id="mobno" name="mobno" required>
              </div>
            </div>
            <div class="form-group">
              <label class="control-label col-sm-2" for="contact">Comment:</label>
              <div class="col-sm-5">
                <textarea   class="form-control" id="contact" name="contact"></textarea>
              </div>
            </div>
            <div class="form-group">
              <label class="control-label col-sm-2" for="contact">Captcha Code:</label>
              <div class="col-sm-3">
                <input type="text" class="form-control" id="captcha" name="captcha" required>
              </div>
              <div class="col-sm-2"><img src="captcha.php" id="captcha_id"></div>
              <div class="col-sm-2">
              <i class="fa fa-refresh" onclick="
                    document.getElementById('captcha_id').src ='captcha.php?'+(new Date());return true;"
               style="font-size:25px;color:red;cursor:pointer"></i>
              </div>
            </div>
            <input type="submit"  name="submit" class="next btn btn-success" 
                value="SUBMIT" id="submit" style='margin-left:30%'/>
          </div>                         
        </form>
      </div>
    </div>
</div> 
</body> 
</html>
