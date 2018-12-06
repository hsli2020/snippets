<?php
use Phalcon\UserPlugin\Models\Location\Locations;

class PlacesController extends BaseController
{
    public $save;

    public function saveAction()
    {
        parent::saveAction();

        $st_toSave = $this->request->getPost();
        $lang = array('language' => $this->_activeLanguage);

        $condition = 'language="'.$this->_activeLanguage.'" and geo_point = POINT("'.$st_toSave['latitude'].'","'.$st_toSave['longitude'].'")';
        $location = Locations::findFirst($condition);

        if($location) {
            $st_output = array('id' => $location->getId());
            $this->response->setJsonContent($st_output);

            $this->initResponse();
            $this->setPayload($st_output);

            return $this->render();
        }

        $st_toSave['geo_point'] = new \Phalcon\Db\RawValue('Point('.$st_toSave['latitude'].', '.$st_toSave['longitude'].')');

        $location = new Locations();
        $location->assign(array_merge($st_toSave, $lang));

        if(false == $location->create())
        {
            $st_output = array('error' => '?');
        }

        $st_output = array('id' => $location->getId());
        $this->response->setJsonContent($st_output);

        $this->initResponse();
        $this->setPayload($st_output);

        return $this->render();
    }
}
