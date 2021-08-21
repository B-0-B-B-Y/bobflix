import { Loader } from '../components/Loader'
import styles from '../styles/Home.module.scss'

export const Home = () => {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Bobflix</h1>
      <Loader />
    </div>
  )
}

export default Home
