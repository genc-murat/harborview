import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Button } from '@/components/ui/button';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Play, Square, RefreshCw, Trash2, AlertCircle } from 'lucide-react';
import { useToast } from '@/components/ui/use-toast';
import { api } from '@/lib/api';

const Containers = () => {
  const [containers, setContainers] = useState([]);
  const [loading, setLoading] = useState(true);
  const { toast } = useToast();

  useEffect(() => {
    fetchContainers();
    // Her 30 saniyede bir güncelle
    const interval = setInterval(fetchContainers, 30000);
    return () => clearInterval(interval);
  }, []);

  const fetchContainers = async () => {
    setLoading(true);
    try {
      const data = await api.getContainers();
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
      switch (action) {
        case 'start':
          await api.startContainer(id);
          break;
        case 'stop':
          await api.stopContainer(id);
          break;
        case 'remove':
          await api.removeContainer(id);
          break;
        default:
          return;
      }
      
      await fetchContainers();
      toast({
        title: 'Success',
        description: `Container ${action} successful`,
      });
    } catch {
      toast({
        title: 'Error',
        description: `Failed to ${action} container`,
        variant: 'destructive',
      });
    }
  };

  const getStatusColor = (status) => {
    if (status.includes('Up')) return 'text-green-600';
    if (status.includes('Exited')) return 'text-red-600';
    return 'text-yellow-600';
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-full">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary" />
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Containers</h1>
        <Button onClick={fetchContainers} disabled={loading}>
          <RefreshCw className={`mr-2 h-4 w-4 ${loading ? 'animate-spin' : ''}`} />
          Refresh
        </Button>
      </div>

      {containers.length === 0 ? (
        <div className="text-center py-12 bg-muted/50 rounded-lg">
          <AlertCircle className="mx-auto h-12 w-12 text-muted-foreground" />
          <h3 className="mt-4 text-lg font-semibold">No Containers Found</h3>
          <p className="text-muted-foreground">Start by running some containers</p>
        </div>
      ) : (
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Image</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Created</TableHead>
              <TableHead className="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {containers.map((container) => (
              <TableRow key={container.Id} className="hover:bg-muted/50">
                <TableCell className="font-medium">
                  <Link 
                    to={`/containers/${container.Id}`} 
                    className="text-primary hover:underline"
                  >
                    {container.Names[0].replace('/', '')}
                  </Link>
                </TableCell>
                <TableCell>{container.Image}</TableCell>
                <TableCell className={getStatusColor(container.Status)}>
                  {container.Status}
                </TableCell>
                <TableCell>{new Date(container.Created * 1000).toLocaleString()}</TableCell>
                <TableCell className="text-right">
                  <div className="flex justify-end space-x-2">
                    <Button
                      variant="ghost"
                      size="icon"
                      onClick={() => handleContainerAction(container.Id, 'start')}
                      className="hover:bg-green-100 hover:text-green-600"
                      disabled={container.Status.includes('Up')}
                    >
                      <Play className="h-4 w-4" />
                    </Button>
                    <Button
                      variant="ghost"
                      size="icon"
                      onClick={() => handleContainerAction(container.Id, 'stop')}
                      className="hover:bg-yellow-100 hover:text-yellow-600"
                      disabled={container.Status.includes('Exited')}
                    >
                      <Square className="h-4 w-4" />
                    </Button>
                    <Button
                      variant="ghost"
                      size="icon"
                      onClick={() => handleContainerAction(container.Id, 'remove')}
                      className="hover:bg-red-100 hover:text-red-600"
                    >
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </div>
  );
};

export default Containers;