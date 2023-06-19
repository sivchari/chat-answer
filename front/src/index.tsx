import { createRoot } from 'react-dom/client';
import App from 'src/components/App.tsx';

const container = document.getElementById('root');
const root = createRoot(container);
root.render(<App />);
