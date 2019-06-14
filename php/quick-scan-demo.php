<?php
session_start();

$loc = '';
$saved = 0;
$scanned = isset($_SESSION["scanned"]) ? $_SESSION["scanned"] : [];

if ($_POST) {
    $btn = $_POST['btn'];

    // user clicked 'Add' button
    if ($btn == 'Add') {
        $_POST['shelf'] = '';
        $loc = $_POST['loc'];
        $scanned[] = $_POST;
        $_SESSION["scanned"] = $scanned;
        $saved = 1;
    }

    // user clicked 'Export' button
    if ($btn == 'Export') {
        $filename = "./_2/quick-scan-demo.csv";
        $fp = fopen($filename, 'w');
        fputcsv($fp, ['LOC', 'QTY', 'UPC']);
        foreach ($scanned as $row) {
            fputcsv($fp, [
                $row['loc'],
                $row['qty'],
                $row['upc'],
            ]);
        }
        fclose($fp);
        $message = "Exported to <b>$filename</b>";
    }

    // user clicked 'Delete Selected' button
    if ($btn == 'Delete Selected') {
        $items = isset($_POST['items']) ? $_POST['items'] : [];
        foreach ($items as $key) {
            array_splice($scanned, $key, 1);
        }
        $_SESSION["scanned"] = $scanned;
    }

    // user clicked 'Delete All' button
    if ($btn == 'Delete All') {
        $scanned = $_SESSION["scanned"] = [];
    }
}

$scanned = array_reverse($scanned);
?>
<html>
<head>
<title>Quick Scan Demo</title>
<style>
  body { width: 960px; margin: 0 auto; }
  table { border-collapse: collapse; width: 100%; }
  table, td, th { border: 1px solid gray; padding: 5px 10px; }
  input[type=submit] { padding: 4px 10px; }
  input[type=text] { font-size:15px; padding: 5px 10px; }
  .center { text-align: center; }
</style>
</head>
<body>
<h2>Quick Scan Demo</h2>
<form name="form1" method="POST">
  <input type="text" id="loc" name="loc" value="<?= $loc ?>" size="15" placeholder="loc">
  <input type="text" id="qty" name="qty" value="1" size="15" placeholder="qty">
  <input type="text" id="upc" name="upc" value="" size="20" placeholder="upc" autofocus>

  <input type="submit" name="btn" value="Add" id="add" />
  <input type="submit" name="btn" value="Export" />
  <input type="submit" name="btn" value="Delete Selected">
  <input type="submit" name="btn" value="Delete All" onclick="return confirm('Are you sure?')">

  <p>
    Number of items: <b style="font-size: 24px;"><?= count($scanned) ?></b>
    <?php if (isset($message)) { ?><span style="float: right"><?= $message ?></span><?php } ?>
  </p>

  <table border="1">
    <tr style="background-color: #eee;">
      <th>#</th>
      <th>LOC</th>
      <th>QTY</th>
      <th>UPC</th>
      <th>SHELF</th>
      <th>SELECT</th>
    </tr>
    <?php foreach ($scanned as $i => $row) { ?>
    <tr>
        <td class="center"><?= count($scanned)-$i ?></td>
        <td><?= $row['loc'] ?></td>
        <td class="center"><?= $row['qty'] ?></td>
        <td><?= $row['upc'] ?></td>
        <td><?= $row['shelf'] ?></td>
        <td class="center"><input type="checkbox"  name="items[]" value="<?= count($scanned)-$i-1 ?>"></td>
    </tr>
    <?php } ?>
  </table>
</form>

<script type="text/javascript">
  var addbtn = document.getElementById("add");
  addbtn.addEventListener("focus", submitForm);
  function submitForm(){
    addbtn.click();
  }

  var saved = <?= $saved ?>;
  if (saved) {
    var audio = new Audio('http://192.168.0.12/assets/sound/sound4.mp3');
    audio.play();
  }
</script>

</body>
</html>
