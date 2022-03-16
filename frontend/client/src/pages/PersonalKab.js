import React from 'react';
import '../style/personalKab.css'

export default function PersonalKab(props) {
    const name = props.name
    
   return (
       <div className='personalKab'>
           <div className='content__personalKab'>
               <p className='mainText__personalKab'>Select a project</p>
               <div className='flexArray__personalKab'>

               </div>
           </div>
       </div>
   )
};

