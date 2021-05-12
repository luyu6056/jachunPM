<?php
/**
 * The edit view of project module of ZenTaoPMS.
 *
 * @copyright   Copyright 2009-2015 青岛易软天创网络科技有限公司(QingDao Nature Easy Soft Network Technology Co,LTD, www.cnezsoft.com)
 * @license     ZPL (http://zpl.pub/page/zplv12.html)
 * @author      Chunsheng Wang <chunsheng@cnezsoft.com>
 * @package     project
 * @version     $Id: edit.html.php 4728 2013-05-03 06:14:34Z chencongzhi520@gmail.com $
 * @link        http://www.zentao.net
 */
?>
<?php include '../../common/view/header.html.php';?>
<?php include '../../common/view/datepicker.html.php';?>
<?php include '../../common/view/kindeditor.html.php';?>
<?php js::import($jsRoot . 'misc/date.js');?>
<div id='mainContent' class='main-content'>
  <div class='center-block'>
    <div class='main-header'>
      <h2>
        <span class='prefix label-id'><strong><?php echo $project->id;?></strong></span>
        <?php echo html::a($this->createLink('project', 'view', 'project=' . $project->id), $project->name, '_blank');?>
        <small><?php echo $lang->arrow . ' ' . $lang->project->edit;?></small>
      </h2>
    </div>
    <form class='load-indicator main-form form-ajax' method='post' target='hiddenwin' id='dataform'>
      <table class='table table-form'>
        <tr>
          <th><?php echo $lang->project->name;?></th>
          <td><?php echo html::input('name', $project->name, "class='form-control' autocomplete='off' required");?></td><td></td>
        </tr>
        <tr>
          <th><?php echo $lang->project->code;?></th>
          <td><?php echo html::input('code', $project->code, "class='form-control' autocomplete='off' required");?></td>
        </tr>
        <tr>
          <th><?php echo $lang->project->dateRange;?></th>
          <td>
            <div class='input-group'>
              <?php echo html::input('begin', $project->begin, "class='form-control form-date' onchange='computeWorkDays()' required placeholder='" . $lang->project->begin . "'");?>
              <span class='input-group-addon fix-border'><?php echo $lang->project->to;?></span>
              <?php echo html::input('end', $project->end, "class='form-control form-date' onchange='computeWorkDays()' required placeholder='" . $lang->project->end . "'");?>
              <div class='input-group-btn'>
                <button type='button' class='btn dropdown-toggle' data-toggle='dropdown'><?php echo $lang->project->byPeriod;?> <span class='caret'></span></button>
                <ul class='dropdown-menu'>
                  <?php foreach ($lang->project->endList as $key => $name):?>
                  <li><a href='javascript:computeEndDate("<?php echo $key;?>")'><?php echo $name;?></a></li>
                  <?php endforeach;?>
                </ul>
              </div>
            </div>
          </td>
        </tr>
        <tr>
          <th><?php echo $lang->project->days;?></th>
          <td>
            <div class='input-group'>
              <?php echo html::input('days', $project->days, "class='form-control' autocomplete='off'");?>
              <span class='input-group-addon'><?php echo $lang->project->day;?></span>
            </div>
          </td>
        </tr>
        <tr>
          <th><?php echo $lang->project->type;?></th>
          <td><?php echo html::select('type', $lang->project->typeList, $project->type, "class='form-control'");?></td>
        </tr>
        <tr>
          <th><?php echo $lang->project->teamname;?></th>
          <td><?php echo html::input('team', $project->team, "class='form-control' autocomplete='off'");?></td>
        </tr>
        <tr>
          <th><?php echo $lang->project->status;?></th>
          <td><?php echo html::select('status', $lang->project->statusList, $project->status, "class='form-control'");?></td>
        </tr>
        <?php if($config->global->flow == 'onlyTask'):?>
        <tr>
          <th><?php echo $lang->project->owner;?></th>
          <td><?php echo html::select('PM', $pmUsers, $project->PM, "class='form-control chosen'");?></td>
        </tr>
        <?php else:?>
        <tr>
          <th rowspan='2'><?php echo $lang->project->owner;?></th>
          <td>
            <div class='input-group'>
              <span class='input-group-addon'><?php echo $lang->project->PO;?></span>
              <?php echo html::select('PO', $poUsers, $project->PO, "class='form-control chosen'");?>
            </div>
          </td>
          <td>
            <div class='input-group'>
              <span class='input-group-addon'><?php echo $lang->project->QD;?></span>
              <?php echo html::select('QD', $qdUsers, $project->QD, "class='form-control chosen'");?>
            </div>
          </td>
        </tr>
        <tr>
          <td>
            <div class='input-group'>
              <span class='input-group-addon'><?php echo $lang->project->PM;?></span>
              <?php echo html::select('PM', $pmUsers, $project->PM, "class='form-control chosen'");?>
            </div>
          </td>
          <td>
            <div class='input-group'>
              <span class='input-group-addon'><?php echo $lang->project->RD;?></span>
              <?php echo html::select('RD', $rdUsers, $project->RD, "class='form-control chosen'");?>
            </div>
          </td>
        </tr>
        <?php endif;?>
        <tr <?php if($config->global->flow == 'onlyTask') echo "class='hidden'";?>>
          <th><?php echo $lang->project->manageProducts;?></th>
          <td class='text-left' id='productsBox' colspan="2">
            <div class='row'>
              <?php $i = 0;?>
              <?php foreach($linkedProducts as $product):?>
              <div class='col-sm-4'>
                <?php $hasBranch = $product->type != 'normal' and isset($branchGroups[$product->id]);?>
                <div class="input-group<?php if($hasBranch) echo ' has-branch';?>">
                  <?php echo html::select("products[$i]", $allProducts, $product->id, "class='form-control chosen' onchange='loadBranches(this)' data-last='" . $product->id . "'");?>
                  <span class='input-group-addon fix-border'></span>
                  <?php if($hasBranch) echo html::select("branch[$i]", $branchGroups[$product->id], $product->branch, "class='form-control chosen' onchange=\"loadPlans('#products{$i}', this.value)\"");?> 
                </div>
              </div>
              <?php $i++;?>
              <?php endforeach;?>
              <div class='col-sm-4'>
                <div class='input-group'>
                  <?php echo html::select("products[$i]", $allProducts, '', "class='form-control chosen' onchange='loadBranches(this)'");?>
                  <span class='input-group-addon fix-border'></span>
                </div>
              </div>
            </div>
          </td>
        </tr>
        <tr <?php if($config->global->flow == 'onlyTask') echo "class='hidden'";?>>
          <th><?php echo $lang->project->linkPlan;?></th>
          <td id="plansBox" colspan="2">
            <div class='row'>
              <?php $i = 0;?>
              <?php foreach($linkedProducts as $product):?>
              <?php $plans = zget($productPlans, $product->id, array(0 => ''));?>
              <div class="col-sm-4" id="plan<?php echo $i;?>"><?php echo html::select("plans[" . $product->id . "]", $plans, $product->plan, "class='form-control chosen'");?></div>
              <?php $i++;?>
              <?php endforeach;?>
            </div>
          </td>
        </tr>
        <tr>
          <th><?php echo $lang->project->desc;?></th>
          <td colspan='2'><?php echo html::textarea('desc', htmlspecialchars($project->desc), "rows='6' class='form-control kindeditor' hidefocus='true'");?></td>
        </tr>
         <tr>
          <th><?php echo $lang->files;?></th>
          <td colspan='3'><?php echo $this->fetch('file', 'buildform', array('fileCount' => '1', 'percent' => '0.9','filesName'=>'files','labelsName'=>'labels','examine'=>'','action'=>'project'));?></td>
        </tr>
        <tr>
          <th><?php echo $lang->project->acl;?></th>
          <td colspan='2'><?php echo nl2br(html::radio('acl', $lang->project->aclList, $project->acl, "onclick='setWhite(this.value);'", 'block'));?></td>
        </tr>
        <tr>
          <th><?php echo $lang->project->ftpPath;?></th>
          <td><?php echo html::input('ftpPath', $project->ftpPath, "class='form-control' autocomplete='off' required");?></td>
        </tr>
        <tr id='whitelistBox' <?php if($project->acl != 'custom') echo "class='hidden'";?>>
          <th><?php echo $lang->project->whitelist;?></th>
          <td colspan='2'><?php echo html::checkbox('whitelist', $groups, $project->whitelist, '', '', 'inline');?></td>
        </tr>
        <tr><td colspan='3' class='text-center form-actions'><?php echo html::submitButton() . ' ' . html::backButton();?></td></tr>
      </table>
    </form>
  </div>
</div>
<?php js::set('weekend', $config->project->weekend);?>
<?php js::set('errorSameProducts', $lang->project->errorSameProducts);?>
<?php include '../../common/view/footer.html.php';?>
