<?php
/**
 * The complete file of task module of ZenTaoPMS.
 *
 * @copyright   Copyright 2009-2015 青岛易软天创网络科技有限公司(QingDao Nature Easy Soft Network Technology Co,LTD, www.cnezsoft.com)
 * @license     ZPL (http://zpl.pub/page/zplv12.html)
 * @author      Jia Fu <fujia@cnezsoft.com>
 * @package     task
 * @version     $Id: complete.html.php 935 2010-07-06 07:49:24Z jajacn@126.com $
 * @link        http://www.zentao.net
 */
?>
<?php include '../../common/view/header.html.php';?>
<?php include '../../common/view/kindeditor.html.php';?>
<?php include '../../common/view/datepicker.html.php';?>
<div id='mainContent' class='main-content'>
  <div class='center-block' style="min-height: 600px;">
   <?php echo $this->fetch('file', 'printFiles', array('files' => $project->file, 'fieldset' => 'true',type=>'spec'));?>
    </div>
  </div>
</div>
<?php include '../../common/view/footer.html.php';?>
