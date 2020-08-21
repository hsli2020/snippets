==================== ./src/index.js

import React    from 'react'
import ReactDOM from 'react-dom'
import App      from './App'
import { BrowserRouter } from "react-router-dom"
import './index.css'

ReactDOM.render(
  <BrowserRouter>
    <App />
  </BrowserRouter>,
  document.getElementById('root')
)

==================== ./src/App.js

import { Routes, Route } from 'react-router-dom'
import Home   from './pages/Home'
import Course from './pages/Course'
import Lesson from './pages/Lesson'

function App() {
  return (
    <div className="App">
      <main>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/courses/:courseId" element={<Course />} />
          <Route path="/courses/:courseId/lessons/:lessonId" element={<Lesson />} />
        </Routes>
      </main>
    </div>
  )
}

export default App

==================== ./src/components/CompleteAndContinueButtons.js

import { useNavigate } from 'react-router-dom'

function CompleteAndContinueButton(props) {
  const navigate = useNavigate()

  function completeAndContinue () {
    navigate(`/courses/${props.courseId}/lessons/${props.lessonId}`)
  }

  return (
    <button className="button primary" onClick={completeAndContinue}>
      Complete and continue
    </button>
  )
}

export default CompleteAndContinueButton

==================== ./src/components/CourseSummary.js

import { Link } from 'react-router-dom'

function CourseSummary(props) {
  return (
    <section key={props.course.id} className="summary">
      <div>
        <div className="title">
          <h2>
            <Link
              className="no-underline cursor-pointer"
              to={'/courses/' + props.course.id}
            >
              {props.course.title}
            </Link>
          </h2>
        </div>
        <p>
          <Link
            className="no-underline cursor-pointer"
            to={'/courses/' + props.course.id}
          >
            {props.course.description}
          </Link>
        </p>
      </div>
    </section>
  )
}

export default CourseSummary

==================== ./src/components/LessonSummary.js

import { Link } from 'react-router-dom'

function LessonSummary(props) {
  return (
    <section key={props.lesson.id} className="summary">
      <div>
        <div className="title">
          <h2>
            <Link
              className="no-underline cursor-pointer"
              to={'/courses/' + props.courseId + '/lessons/' + props.lesson.id}
            >
              {props.num}. {props.lesson.title}
            </Link>
          </h2>
        </div>
        <p>
          <Link
            className="no-underline cursor-pointer"
            to={'/courses/' + props.courseId + '/lessons/' + props.lesson.id}
          >
            {props.lesson.description}
          </Link>
        </p>
      </div>
    </section>
  )
}

export default LessonSummary

==================== ./src/courses.js

const courses = [
  {
    id: 1,
    title: "Photography for Beginners",
    description: "Phasellus ac tellus tincidunt, pharetra dui eu, bibendum nulla.",
    lessons: [
      {
        id: 1,
        title: "Welcome to the course",
        description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. ...",
        vimeoId: 76979871
      },
      {
        id: 2,
        title: "How does a camera work?",
        description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. ...",
        vimeoId: 76979871
      },
    ]
  },
  {
    id: 2,
    title: "Advanced Photography",
    description: "Cras ut sem eu ligula luctus ornare quis nec arcu.",
    lessons: [...]
  }
]

export default courses

==================== ./src/pages/Course.js

import { useParams } from 'react-router-dom'
import LessonSummary from '../components/LessonSummary'
import { Link } from 'react-router-dom'
import courses from '../courses'

function Course() {
  const { courseId } = useParams()
  const course = courses.find(course => course.id === parseInt(courseId))

  return (
    <div className="Course page">
      <header>
        <p>
          <Link to={'/'}>Back to courses</Link>
        </p>
        <h1>{course.title}</h1>
        <p>{course.description}</p>
        <Link className="button primary icon" 
            to={`/courses/${courseId}/lessons/${course.lessons[0].id}`}>
          Start course
        </Link>
      </header>
      <div>
        {course.lessons.map((lesson, index) => (
          <LessonSummary
            courseId={courseId}
            lesson={lesson}
            num={index + 1}
            key={lesson.id}
          />
        ))}
      </div>
    </div>
  )
}

export default Course

==================== ./src/pages/Home.js

import CourseSummary from '../components/CourseSummary'
import courses from '../courses'

function Home() {
  return (
    <div className="Home page">
      <header>
        <h1>React Online Course Site</h1>
      </header>
      {courses.map((course) => (
        <CourseSummary course={course} key={course.id} />
      ))}
    </div>
  )
}

export default Home

==================== ./src/pages/Lesson.js

import { Link, useParams } from 'react-router-dom'
import Vimeo from '@u-wave/react-vimeo'
import CompleteAndContinueButton from "../components/CompleteAndContinueButtons";
import courses from '../courses'

function Lesson() {
  const { courseId, lessonId } = useParams()

  const course = courses.find(course => course.id === parseInt(courseId))
  const lesson = course.lessons.find(lesson => lesson.id === parseInt(lessonId))

  const nextLessonId = () => {
    const currentIndex = course.lessons.indexOf(lesson)
    const nextIndex = (currentIndex + 1) % course.lessons.length
    return course.lessons[nextIndex].id
  }

  return (
    <div className="Lesson page">
      <header>
        <p>
          <Link to={'/courses/' + course.id}>Back to {course.title}</Link>
        </p>
        <h1>{lesson.title}</h1>
      </header>
      <div className="Content">
        <Vimeo video={lesson.vimeoId} responsive />
        <CompleteAndContinueButton courseId={courseId} lessonId={nextLessonId()}/>
      </div>
    </div>
  )
}

export default Lesson
