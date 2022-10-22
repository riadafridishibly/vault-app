import Generator from '../../components/Generator/Generator'
import styles from './PasswordGenerator.module.css'

function PasswordGenerator() {
    return (
        <div className={styles.container}>
            <Generator />
        </div>
    )
}

export default PasswordGenerator