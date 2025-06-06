// WUT - Professional Wu-Tang Documentation JavaScript

document.addEventListener('DOMContentLoaded', function() {
    // Smooth scrolling for navigation links
    const navLinks = document.querySelectorAll('.nav-link[href^="#"]');
    navLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            const target = document.querySelector(this.getAttribute('href'));
            if (target) {
                target.scrollIntoView({
                    behavior: 'smooth',
                    block: 'start'
                });
            }
        });
    });

    // Add scroll effect to navbar
    const navbar = document.querySelector('.navbar');
    let lastScrollTop = 0;
    
    window.addEventListener('scroll', function() {
        const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
        
        if (scrollTop > 100) {
            navbar.style.background = 'rgba(255, 255, 255, 0.95)';
            navbar.style.backdropFilter = 'blur(10px)';
        } else {
            navbar.style.background = 'var(--bg-primary)';
            navbar.style.backdropFilter = 'blur(8px)';
        }
        
        lastScrollTop = scrollTop;
    });

    // Add intersection observer for animations
    const observerOptions = {
        threshold: 0.1,
        rootMargin: '0px 0px -50px 0px'
    };

    const observer = new IntersectionObserver(function(entries) {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.style.opacity = '1';
                entry.target.style.transform = 'translateY(0)';
            }
        });
    }, observerOptions);

    // Animate sections on scroll
    const animateElements = document.querySelectorAll('.feature-card, .package-card, .example-card, .install-method');
    animateElements.forEach(el => {
        el.style.opacity = '0';
        el.style.transform = 'translateY(20px)';
        el.style.transition = 'opacity 0.6s ease, transform 0.6s ease';
        observer.observe(el);
    });

    // Add Wu-Tang wisdom tooltip system
    const wisdomQuotes = [
        "Time is the most valuable currency, God",
        "Protect ya neck... and ya schedule",
        "Cash rules everything around me, including my calendar",
        "Wu-Tang is for the children... and productivity",
        "Diversify yo bonds... and time management strategies",
        "Bring da ruckus to your daily routine",
        "36 chambers of supreme focus"
    ];

    // Add random Wu-Tang wisdom on logo hover
    const logo = document.querySelector('.logo');
    if (logo) {
        logo.addEventListener('mouseenter', function() {
            const randomWisdom = wisdomQuotes[Math.floor(Math.random() * wisdomQuotes.length)];
            
            // Create tooltip
            const tooltip = document.createElement('div');
            tooltip.className = 'wu-tooltip';
            tooltip.textContent = randomWisdom;
            tooltip.style.cssText = `
                position: absolute;
                background: var(--bg-dark);
                color: var(--primary-color);
                padding: 8px 12px;
                border-radius: 6px;
                font-size: 12px;
                white-space: nowrap;
                z-index: 1000;
                top: 100%;
                left: 50%;
                transform: translateX(-50%);
                margin-top: 8px;
                box-shadow: var(--shadow);
                font-family: var(--font-mono);
            `;
            
            this.style.position = 'relative';
            this.appendChild(tooltip);
            
            // Remove tooltip after 3 seconds
            setTimeout(() => {
                if (tooltip.parentNode) {
                    tooltip.parentNode.removeChild(tooltip);
                }
            }, 3000);
        });
    }

    // Add Wu-Tang mathematics to version badge
    const version = document.querySelector('.version');
    if (version) {
        version.addEventListener('click', function() {
            const mathematics = ['Knowledge', 'Wisdom', 'Understanding', 'Culture/Freedom', 'Power/Refinement'];
            const randomMath = mathematics[Math.floor(Math.random() * mathematics.length)];
            const originalText = this.textContent;
            
            this.textContent = randomMath;
            this.style.background = 'var(--secondary-color)';
            
            setTimeout(() => {
                this.textContent = originalText;
                this.style.background = 'var(--primary-color)';
            }, 2000);
        });
    }

    // Add konami code for Wu-Tang easter egg
    let konamiCode = [];
    const konamiSequence = [
        'ArrowUp', 'ArrowUp', 'ArrowDown', 'ArrowDown',
        'ArrowLeft', 'ArrowRight', 'ArrowLeft', 'ArrowRight',
        'KeyB', 'KeyA'
    ];

    document.addEventListener('keydown', function(e) {
        konamiCode.push(e.code);
        
        if (konamiCode.length > konamiSequence.length) {
            konamiCode.shift();
        }
        
        if (konamiCode.length === konamiSequence.length) {
            const match = konamiCode.every((code, index) => code === konamiSequence[index]);
            
            if (match) {
                // Wu-Tang easter egg
                const body = document.body;
                body.style.filter = 'hue-rotate(36deg) saturate(1.5)';
                body.style.animation = 'wu-chaos 2s ease-in-out';
                
                // Add CSS animation if not exists
                if (!document.querySelector('#wu-chaos-style')) {
                    const style = document.createElement('style');
                    style.id = 'wu-chaos-style';
                    style.textContent = `
                        @keyframes wu-chaos {
                            0%, 100% { transform: scale(1) rotate(0deg); }
                            25% { transform: scale(1.02) rotate(1deg); }
                            50% { transform: scale(0.98) rotate(-1deg); }
                            75% { transform: scale(1.01) rotate(0.5deg); }
                        }
                    `;
                    document.head.appendChild(style);
                }
                
                // Show Wu-Tang message
                const message = document.createElement('div');
                message.textContent = '🐉 WU-TANG CLAN AIN\'T NUTHIN\' TA F\' WIT! 🐉';
                message.style.cssText = `
                    position: fixed;
                    top: 50%;
                    left: 50%;
                    transform: translate(-50%, -50%);
                    background: var(--primary-color);
                    color: var(--bg-dark);
                    padding: 20px 40px;
                    border-radius: 12px;
                    font-weight: bold;
                    font-size: 24px;
                    z-index: 10000;
                    box-shadow: var(--shadow-hover);
                    animation: wu-pulse 0.5s ease-in-out;
                `;
                
                document.body.appendChild(message);
                
                setTimeout(() => {
                    body.style.filter = '';
                    body.style.animation = '';
                    if (message.parentNode) {
                        message.parentNode.removeChild(message);
                    }
                }, 3000);
                
                konamiCode = [];
            }
        }
    });

    // Add loading animation for demo gif
    const demoGif = document.querySelector('.demo-gif');
    if (demoGif) {
        demoGif.addEventListener('load', function() {
            this.style.opacity = '1';
            this.style.transform = 'scale(1)';
        });
        
        // Initial state
        demoGif.style.opacity = '0';
        demoGif.style.transform = 'scale(0.95)';
        demoGif.style.transition = 'opacity 0.5s ease, transform 0.5s ease';
    }
});