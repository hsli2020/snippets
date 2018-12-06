<?php

// http://paul-m-jones.com/archives/6043
// http://pmjones.io/adr/

class BlogController
{
    // POST /blog/{id}
    public function update($id)
    {
        $blog = $this->model->fetch($id);
        if (! $blog) {
            // 404 Not Found
            // (no blog entry with that ID)
            $this->response->status->set(404);
            $this->view->setData(array('id' => $id));
            $content = $this->view->render('not-found');
            $this->response->body->setContent($content);
            return;
        }

        $data = $this->request->post->get('blog');
        if (! $blog->update($data)) {
            // update failure, but why?
            if (! $blog->isValid()) {
                // 422 Unprocessable Entity
                // (not valid)
                $this->response->status->set(422);
                $this->view->setData(array('blog' => $blog));
                $content = $this->view->render('update');
                $this->response->body->setContent($content);
                return;
            } else {
                // 500 Server Error
                // (i.e., valid data, but update failed for some other reason)
                $this->response->status->set(500);
                return;
            }
        }

        // 200 OK
        // (i.e., the update worked)
        $this->response->status->set(200);
        $this->view->setData(array('blog' => $blog));
        $content = $this->view->render('update');
        $this->response->body->setContent($content);
    }
}

class BlogUpdateAction
{
    // POST /blog/{id}
    public function __invoke($id)
    {
        $data = $this->request->post->get('blog');
        $blog = $this->domain->update($id, $data);
        $this->responder->setData('id' => $id, 'blog' => $blog);
        $this->responder->__invoke();
    }
}

class BlogUpdateResponder
{
    public function __invoke()
    {
        if (! $this->data->blog) {
            // 404 Not Found
            // (no blog entry with that ID)
            $this->response->setStatus(404);
            $this->view->setData($this->data);
            $content = $this->view->render('not-found');
            $this->response->body->setContent($content);
            return;
        }

        if ($this->data->blog->updateFailed()) {
            // 500 Server Error
            // (i.e., valid data, but update failed for some other reason)
            $this->response->status->set(500);
            return;
        }

        if (! $this->data->blog->isValid()) {
            // 422 Unprocessable Entity
            // (invalid data submitted)
            $this->response->setStatus(422);
            $this->view->setData($this->data);
            $content = $this->view->render('update');
            $this->response->body->setContent($content);
            return;
        }

        // 200 OK
        // (i.e., the update worked)
        $this->view->setData($this->data);
        $content = $this->view->render('update');
        $this->response->body->setContent($content);
    }
}

