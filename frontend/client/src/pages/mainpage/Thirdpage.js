import '../../style/pics.css';
import '../../style/App.css';
import ScrollButton from '../../style/ScrollTop';
import Register from '../RegForm';

function Thirdpage() {
    return (
        <>
        <div className='thirdpagebg'>
            <div className="thirdpage" id="thirdpage">
                <div className='notHeader'>
                    <ScrollButton />
                    </div>
            <Register />
                
            </div>
            </div>
        </>
    );
}

export default Thirdpage;