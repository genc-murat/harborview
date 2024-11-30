import { NavLink } from 'react-router-dom';
import { Box, Database, LayoutDashboard } from 'lucide-react';

const Sidebar = () => {
  const navLinkClass = ({ isActive }) =>
    `group flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors ${
      isActive 
        ? 'bg-primary-light text-primary' 
        : 'text-muted-foreground hover:bg-accent hover:text-accent-foreground'
    }`;

  return (
    <div className="hidden md:flex md:w-64 md:flex-col">
      <div className="flex flex-col flex-grow pt-5 bg-card border-r overflow-y-auto">
        <div className="flex items-center flex-shrink-0 px-4 mb-5">
          <h1 className="text-xl font-semibold text-primary">Docker Manager</h1>
        </div>
        <nav className="flex-1 px-2 space-y-1">
          <NavLink to="/" className={navLinkClass} end>
            <LayoutDashboard className="mr-3 h-5 w-5" />
            Dashboard
          </NavLink>
          <NavLink to="/containers" className={navLinkClass}>
            <Box className="mr-3 h-5 w-5" />
            Containers
          </NavLink>
          <NavLink to="/images" className={navLinkClass}>
            <Database className="mr-3 h-5 w-5" />
            Images
          </NavLink>
        </nav>
      </div>
    </div>
  );
};

export default Sidebar;