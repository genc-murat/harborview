import { useEffect, useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Box, Play, HardDrive, Cpu } from 'lucide-react';
import { api } from '@/lib/api';
import { cn } from '@/lib/utils';

const Dashboard = () => {
  const [stats, setStats] = useState({
    totalContainers: 0,
    runningContainers: 0,
    totalImages: 0,
    cpuUsage: 0,
    memoryUsage: 0,
    diskUsage: 0
  });

  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        // Container stats
        const containers = await api.getContainers();
        const runningContainers = containers.filter(c => c.State === 'running').length;
        
        // Image stats
        const images = await api.getImages();

        // System stats - örnek veri
        const systemStats = {
          cpuUsage: Math.floor(Math.random() * 100),
          memoryUsage: Math.floor(Math.random() * 100),
          diskUsage: Math.floor(Math.random() * 100)
        };

        setStats({
          totalContainers: containers.length,
          runningContainers,
          totalImages: images.length,
          ...systemStats
        });
      } catch (error) {
        console.error('Failed to fetch stats:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchStats();
    // 30 saniyede bir güncelle
    const interval = setInterval(fetchStats, 30000);
    return () => clearInterval(interval);
  }, []);

  const statCards = [
    {
      title: "Total Containers",
      value: stats.totalContainers,
      icon: <Box className="h-6 w-6 text-primary" />,
      description: "Total number of containers"
    },
    {
      title: "Running Containers",
      value: stats.runningContainers,
      icon: <Play className="h-6 w-6 text-green-500" />,
      valueClass: "text-green-500",
      description: "Currently active containers"
    },
    {
      title: "Total Images",
      value: stats.totalImages,
      icon: <HardDrive className="h-6 w-6 text-blue-500" />,
      valueClass: "text-blue-500",
      description: "Available docker images"
    },
    {
      title: "CPU Usage",
      value: `${stats.cpuUsage}%`,
      icon: <Cpu className="h-6 w-6 text-purple-500" />,
      valueClass: "text-purple-500",
      description: "Current CPU utilization"
    }
  ];

  if (loading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary"></div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h2 className="text-3xl font-bold tracking-tight">Dashboard</h2>
        <div className="text-sm text-muted-foreground">
          Last updated: {new Date().toLocaleTimeString()}
        </div>
      </div>

      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-4">
        {statCards.map((card, index) => (
          <Card key={index} className="bg-card hover:shadow-md transition-all">
            <CardHeader className="flex flex-row items-center justify-between pb-2">
              <CardTitle className="text-sm font-medium text-muted-foreground">
                {card.title}
              </CardTitle>
              {card.icon}
            </CardHeader>
            <CardContent>
              <div className="space-y-1">
                <p className={cn("text-2xl font-bold", card.valueClass)}>
                  {card.value}
                </p>
                <p className="text-xs text-muted-foreground">
                  {card.description}
                </p>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      <div className="grid gap-6 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>System Resources</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-4">
              <div className="space-y-2">
                <div className="flex items-center justify-between">
                  <span className="text-sm font-medium">Memory Usage</span>
                  <span className="text-sm text-muted-foreground">{stats.memoryUsage}%</span>
                </div>
                <div className="h-2 bg-muted rounded">
                  <div 
                    className="h-2 bg-primary rounded" 
                    style={{ width: `${stats.memoryUsage}%` }}
                  />
                </div>
              </div>
              <div className="space-y-2">
                <div className="flex items-center justify-between">
                  <span className="text-sm font-medium">Disk Usage</span>
                  <span className="text-sm text-muted-foreground">{stats.diskUsage}%</span>
                </div>
                <div className="h-2 bg-muted rounded">
                  <div 
                    className="h-2 bg-blue-500 rounded" 
                    style={{ width: `${stats.diskUsage}%` }}
                  />
                </div>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Recent Activity</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-4">
              {/* Simüle edilmiş aktivite listesi */}
              {[1, 2, 3].map((_, i) => (
                <div key={i} className="flex items-center space-x-4">
                  <div className="h-2 w-2 rounded-full bg-primary" />
                  <div className="flex-1">
                    <p className="text-sm">Container nginx-proxy was started</p>
                    <p className="text-xs text-muted-foreground">
                      {new Date(Date.now() - i * 1000 * 60 * 30).toLocaleTimeString()}
                    </p>
                  </div>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default Dashboard;