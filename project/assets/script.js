document.addEventListener('mousemove', (event) => {  
    const eye = document.querySelector('.eye');  
    const pupil = document.querySelector('.pupil');  

    const eyeRect = eye.getBoundingClientRect();  
    const eyeCenterX = eyeRect.left + eyeRect.width / 2;  
    const eyeCenterY = eyeRect.top + eyeRect.height / 2;  
    
    const deltaX = event.clientX - eyeCenterX;  
    const deltaY = event.clientY - eyeCenterY;  
    
    const angle = Math.atan2(deltaY, deltaX);  
    const distance = Math.min(eyeRect.width / 4, Math.sqrt(deltaX * deltaX + deltaY * deltaY));  

    const pupilX = distance * Math.cos(angle);  
    const pupilY = distance * Math.sin(angle);  
    
    pupil.style.transform = `translate(${pupilX}px, ${pupilY}px)`;  
    });
