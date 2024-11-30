import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Button } from '@/components/ui/button';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Play, Square, RefreshCw, Trash2 } from 'lucide-react';
import { useToast } from '@/components/ui/use-toast';

const Containers = () => {
  const [containers, setContainers] = useState([]);
  const [loading, setLoading] = useState(true);
  const { toast } = useToast();

  useEffect(() => {
    fetchContainers();
  }, []);

  const fetchContainers = async () => {
    setLoading(true);
    try {
      const response = await fetch('/containers');
      const data = await response.json();
      setContainers(data);
    } catch {
      toast({
        title: 'Error',
        description: 'Failed to fetch containers',
        variant: 'destructive',
      });
    } finally {
      setLoading(false);
    }
  };

  const handleContainerAction = async (id, action) => {
    try {
      const response = await fetch(`/containers/${id}/${action}`, {
        method: 'POST',
      });
      if (response.ok) {
        fetchContainers();
        toast({
          title: 'Success',
          description: `Container ${action} successful`,
        });
      }
    } catch {
      toast({
        title: 'Error',
        description: `Failed to ${action} container`,
        variant: 'destructive',
      });
    }
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-full">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900" />
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Containers</h1>
        <Button onClick={fetchContainers}>
          <RefreshCw className="mr-2 h-4 w-4" />
          Refresh
        </Button>
      </div>

      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Image</TableHead>
            <TableHead>Status</TableHead>
            <TableHead>Created</TableHead>
            <TableHead>Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {containers.map((container) => (
            <TableRow key={container.Id}>
              <TableCell>
                <Link to={`/containers/${container.Id}`} className="text-blue-600 hover:underline">
                  {container.Names[0]}
                </Link>
              </TableCell>
              <TableCell>{container.Image}</TableCell>
              <TableCell>{container.Status}</TableCell>
              <TableCell>{new Date(container.Created * 1000).toLocaleString()}</TableCell>
              <TableCell>
                <div className="flex space-x-2">
                  <Button
                    variant="ghost"
                    size="icon"
                    onClick={() => handleContainerAction(container.Id, 'start')}
                  >
                    <Play className="h-4 w-4" />
                  </Button>
                  <Button
                    variant="ghost"
                    size="icon"
                    onClick={() => handleContainerAction(container.Id, 'stop')}
                  >
                    <Square className="h-4 w-4" />
                  </Button>
                  <Button
                    variant="ghost"
                    size="icon"
                    onClick={() => handleContainerAction(container.Id, 'remove')}
                  >
                    <Trash2 className="h-4 w-4" />
                  </Button>
                </div>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};

export default Containers;