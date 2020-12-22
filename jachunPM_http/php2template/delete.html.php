<?php
/**
 * The delete view file of user module of ZenTaoPMS.
 *
 * @copyright   Copyright 2009-2015 青岛易软天创网络科技有限公司(QingDao Nature Easy Soft Network Technology Co,LTD, www.cnezsoft.com)
 * @license     ZPL (http://zpl.pub/page/zplv12.html)
 * @author      Yidong Wang <yidong@cnezsoft.com>
 * @package     user
 * @version     $Id$
 * @link        http://www.zentao.net
 */
?>
<?php include '../../common/view/header.lite.html.php';?>
<?php js::import($jsRoot . 'md5.js');?>
<div id='mainContent' class='main-content'>
  <div class='main-header'>
    <h2><?php echo $lang->user->delete;?></h2>
  </div>
  <form method='post' id='dataform' target='hiddenwin' style='padding: 20px 5% 40px'>
    <table class='w-p100 table-form'>
      <tr>
        <th class='w-120px text-right'>
          <?php echo $lang->user->verifyPassword;?>
        </th>
        <td>
          <div class="required required-wrapper"></div>
          <?php echo html::password('verifyPassword', '', "class='form-control disabled-ie-placeholder' placeholder='{$lang->user->placeholder->verify}'");?>
        </td>
      </tr>
      <tr>
        <th></th>
        <td><?php echo html::submitButton($lang->delete);?></td>
      </tr>
    </table>
  </form>
</div>
<?php echo html::hidden('verifyRand', $rand);?>
<?php include '../../common/view/footer.lite.html.php';?>
