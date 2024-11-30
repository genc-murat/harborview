import { Button } from '@/components/ui/button';
import { Bell, Settings } from 'lucide-react';

const Header = () => {
  return (
    <header className="bg-white border-b">
      <div className="flex items-center justify-between p-4">
        <div className="flex items-center space-x-4">
          <Button variant="ghost" size="icon">
            <Bell className="h-5 w-5" />
          </Button>
        </div>
        <div className="flex items-center space-x-4">
          <Button variant="ghost" size="icon">
            <Settings className="h-5 w-5" />
          </Button>
        </div>
      </div>
    </header>
  );
};

export default Header;