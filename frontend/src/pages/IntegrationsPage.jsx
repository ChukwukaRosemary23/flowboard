import { Link } from 'react-router-dom';
import Navbar from '../components/Navbar';
import { Search, ExternalLink } from 'lucide-react';
import { useState } from 'react';

function IntegrationsPage() {
  const [searchTerm, setSearchTerm] = useState('');

  const integrations = [
    { 
      name: "Slack", 
      category: "Communication", 
      icon: "💬", 
      description: "Send messages and notifications to Slack channels", 
      color: "from-purple-500 to-pink-500",
      url: "https://slack.com"
    },
    { 
      name: "Google Drive", 
      category: "Storage", 
      icon: "📁", 
      description: "Attach files from Google Drive to cards", 
      color: "from-blue-500 to-green-500",
      url: "https://drive.google.com"
    },
    { 
      name: "GitHub", 
      category: "Development", 
      icon: "🐙", 
      description: "Link pull requests and issues to cards", 
      color: "from-gray-700 to-gray-900",
      url: "https://github.com"
    },
    { 
      name: "Jira", 
      category: "Project Management", 
      icon: "🎯", 
      description: "Sync tickets between Jira and FlowBoard", 
      color: "from-blue-600 to-blue-700",
      url: "https://www.atlassian.com/software/jira"
    },
    { 
      name: "Gmail", 
      category: "Email", 
      icon: "✉️", 
      description: "Turn emails into tasks automatically", 
      color: "from-red-500 to-red-600",
      url: "https://mail.google.com"
    },
    { 
      name: "Zoom", 
      category: "Video", 
      icon: "📹", 
      description: "Start meetings directly from cards", 
      color: "from-blue-500 to-blue-600",
      url: "https://zoom.us"
    },
    { 
      name: "Dropbox", 
      category: "Storage", 
      icon: "📦", 
      description: "Link Dropbox files to your boards", 
      color: "from-blue-400 to-blue-600",
      url: "https://www.dropbox.com"
    },
    { 
      name: "Microsoft Teams", 
      category: "Communication", 
      icon: "👥", 
      description: "Collaborate with Teams integration", 
      color: "from-purple-600 to-blue-600",
      url: "https://www.microsoft.com/en/microsoft-teams"
    },
    { 
      name: "Figma", 
      category: "Design", 
      icon: "🎨", 
      description: "Embed design files in cards", 
      color: "from-pink-500 to-purple-500",
      url: "https://www.figma.com"
    },
    { 
      name: "Salesforce", 
      category: "CRM", 
      icon: "☁️", 
      description: "Sync customer data with boards", 
      color: "from-blue-400 to-cyan-500",
      url: "https://www.salesforce.com"
    },
    { 
      name: "Notion", 
      category: "Documentation", 
      icon: "📝", 
      description: "Link Notion pages to cards", 
      color: "from-gray-800 to-black",
      url: "https://www.notion.so"
    },
    { 
      name: "Asana", 
      category: "Project Management", 
      icon: "🎯", 
      description: "Import tasks from Asana", 
      color: "from-pink-500 to-red-500",
      url: "https://asana.com"
    },
    { 
      name: "Trello", 
      category: "Migration", 
      icon: "📋", 
      description: "Import boards from Trello", 
      color: "from-blue-500 to-blue-700",
      url: "https://trello.com"
    },
    { 
      name: "Discord", 
      category: "Communication", 
      icon: "🎮", 
      description: "Send updates to Discord servers", 
      color: "from-indigo-500 to-purple-600",
      url: "https://discord.com"
    },
    { 
      name: "Zapier", 
      category: "Automation", 
      icon: "⚡", 
      description: "Connect 5000+ apps via Zapier", 
      color: "from-orange-500 to-red-500",
      url: "https://zapier.com"
    },
    { 
      name: "Google Calendar", 
      category: "Calendar", 
      icon: "📅", 
      description: "Sync due dates with your calendar", 
      color: "from-blue-500 to-green-500",
      url: "https://calendar.google.com"
    },
  ];

  const filteredIntegrations = integrations.filter(int =>
    int.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    int.category.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />

      {/* Hero */}
      <section className="bg-gradient-to-br from-indigo-600 to-purple-600 py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h1 className="text-5xl font-bold text-white mb-6">Work smarter with powerful integrations</h1>
          <p className="text-xl text-indigo-100 max-w-3xl mx-auto mb-8">
            Connect FlowBoard with the tools you already use. Over 100+ integrations available.
          </p>
          
          {/* Search Bar */}
          <div className="max-w-2xl mx-auto">
            <div className="relative">
              <Search className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
              <input
                type="text"
                placeholder="Search integrations..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="w-full pl-12 pr-4 py-4 rounded-lg text-lg focus:outline-none focus:ring-4 focus:ring-white/20"
              />
            </div>
          </div>
        </div>
      </section>

      {/* Integrations Grid */}
      <section className="py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid md:grid-cols-4 gap-6">
            {filteredIntegrations.map((integration, idx) => (
              <IntegrationCard key={idx} integration={integration} />
            ))}
          </div>

          {filteredIntegrations.length === 0 && (
            <div className="text-center py-20">
              <p className="text-gray-600 text-xl">No integrations found. Try a different search term.</p>
            </div>
          )}
        </div>
      </section>

      {/* CTA */}
      <section className="bg-gradient-to-r from-indigo-600 to-purple-600 py-20">
        <div className="max-w-4xl mx-auto px-4 text-center">
          <h2 className="text-4xl font-bold text-white mb-6">Can't find what you need?</h2>
          <p className="text-xl text-indigo-100 mb-8">Build your own integration with our API</p>
          <Link
            to="/register"
            className="inline-block bg-white text-indigo-600 px-8 py-4 rounded-lg font-semibold hover:bg-gray-100 transition shadow-lg"
          >
            Get started →
          </Link>
        </div>
      </section>
    </div>
  );
}

// Integration Card Component - NOW CLICKABLE!
function IntegrationCard({ integration }) {
  const handleClick = () => {
    // Open the integration's website in a new tab
    window.open(integration.url, '_blank', 'noopener,noreferrer');
  };

  return (
    <div
      onClick={handleClick}
      className="bg-white rounded-xl p-6 shadow-md hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-2 cursor-pointer border-2 border-gray-100 hover:border-indigo-500 group"
    >
      <div className={`bg-gradient-to-r ${integration.color} w-16 h-16 rounded-xl flex items-center justify-center mb-4 mx-auto text-3xl group-hover:scale-110 transition-transform`}>
        {integration.icon}
      </div>
      <h3 className="text-center font-bold text-gray-900 mb-2">{integration.name}</h3>
      <p className="text-center text-sm text-gray-500 mb-3">{integration.category}</p>
      <p className="text-center text-sm text-gray-600 mb-4">{integration.description}</p>
      
      {/* Connect Button */}
      <button 
        className="w-full bg-gradient-to-r from-indigo-600 to-purple-600 text-white py-2 rounded-lg text-sm font-semibold hover:from-indigo-700 hover:to-purple-700 transition flex items-center justify-center gap-2 group-hover:shadow-lg"
        onClick={(e) => {
          e.stopPropagation(); // Prevent double trigger
          handleClick();
        }}
      >
        Visit Site <ExternalLink className="w-4 h-4" />
      </button>
    </div>
  );
}

export default IntegrationsPage;