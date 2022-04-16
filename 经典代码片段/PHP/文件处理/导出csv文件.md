[TOC]

## 生成CSV文件并下载
```php
/**
 * 数据导出到excel(csv文件)
 * @param $fileName , 导出的文件名称
 * @param array $tileArray， 一维数组
 * @param array $dataArray，二维数组
 */
function exportCsv($fileName, $tileArray = [], $dataArray = [])
{
    ini_set('memory_limit', '512M');
    ini_set('max_execution_time', 0);
    ob_end_clean();
    ob_start();
    header("Content-Type: text/csv");
    header("Content-Disposition:filename=" . $fileName);
    $fp = fopen('php://output', 'w');
    fwrite($fp, chr(0xEF) . chr(0xBB) . chr(0xBF));// 转码 防止乱码(比如微信昵称)
    fputcsv($fp, $tileArray);
    $index = 0;
    foreach ($dataArray as $item) {
        if ($index == 1000) {
            $index = 0;
            ob_flush();
            flush();
        }
        $index++;
        fputcsv($fp, $item);
    }
    ob_flush();
    flush();
    ob_end_clean();
}

exportCsv(["编号", "名称", "年龄"], [
  [1, "张三", 12],
  [2, "张三", 12],
  [3, "张三", 12],
]);
```