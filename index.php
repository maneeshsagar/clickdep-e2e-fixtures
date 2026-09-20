<?php
if ($_SERVER['REQUEST_URI']==='/log') { error_log('E2E-LOG-MARKER'); }
$env=[]; foreach (getenv() as $k=>$v) if (strpos($k,'E2E_')===0) $env[$k]=$v;
header('content-type: application/json'); echo json_encode(['lang'=>'php','env'=>$env]);
