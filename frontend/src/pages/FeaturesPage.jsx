import { Link } from 'react-router-dom';
import Navbar from '../components/Navbar';
import { Bell, Calendar, Zap, Layout, CheckCircle, Users, ArrowRight } from 'lucide-react';

function FeaturesPage() {
  return (
    <div className="min-h-screen bg-white">
      <Navbar />

      {/* Hero */}
      <section className="bg-gradient-to-br from-blue-600 to-purple-600 py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h1 className="text-5xl font-bold text-white mb-6">Powerful features for modern teams</h1>
          <p className="text-xl text-blue-100 max-w-3xl mx-auto">
            Everything you need to organize work and get things done
          </p>
        </div>
      </section>

      {/* Features Grid */}
      <section className="py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid md:grid-cols-2 gap-16">
            
            {/* Inbox */}
            <FeatureDetail
              icon={<Bell className="w-12 h-12" />}
              title="Inbox"
              description="Capture every idea that comes to mind. Emails, Slack messages, or quick thoughts - everything goes into your Inbox automatically."
              features={[
                "Email to Inbox integration",
                "Slack integration",
                "Quick capture from anywhere",
                "Smart categorization"
              ]}
              color="from-blue-500 to-blue-600"
              link="/features/inbox"
            />

            {/* Planner */}
            <FeatureDetail
              icon={<Calendar className="w-12 h-12" />}
              title="Planner"
              description="Sync your calendar and see your entire day at a glance. Time-block your tasks and stay on schedule."
              features={[
                "Calendar synchronization",
                "Time blocking",
                "Daily/weekly views",
                "Task scheduling"
              ]}
              color="from-purple-500 to-purple-600"
              link="/features/planner"
            />

            {/* Automation */}
            <FeatureDetail
              icon={<Zap className="w-12 h-12" />}
              title="Automation"
              description="No-code automation built into every board. Let Butler handle the repetitive work so you can focus on what matters."
              features={[
                "No-code automation rules",
                "Card automation",
                "Due date reminders",
                "Custom triggers"
              ]}
              color="from-yellow-500 to-yellow-600"
            />

            {/* Power-Ups */}
            <FeatureDetail
              icon={<Layout className="w-12 h-12" />}
              title="Power-Ups"
              description="Extend FlowBoard with integrations to the tools you already use. Connect your entire workflow in one place."
              features={[
                "100+ integrations",
                "Custom Power-Ups",
                "API access",
                "Webhook support"
              ]}
              color="from-green-500 to-green-600"
            />

            {/* Templates */}
            <FeatureDetail
              icon={<CheckCircle className="w-12 h-12" />}
              title="Templates"
              description="Start fast with pre-built templates. From project management to product launches, we've got you covered."
              features={[
                "100+ ready templates",
                "Industry-specific boards",
                "Custom templates",
                "Team templates"
              ]}
              color="from-pink-500 to-pink-600"
            />

            {/* Integrations */}
            <FeatureDetail
              icon={<Users className="w-12 h-12" />}
              title="Integrations"
              description="Connect the apps your team already uses. Slack, Google Drive, GitHub, and hundreds more."
              features={[
                "Popular app integrations",
                "Real-time sync",
                "Two-way updates",
                "Custom webhooks"
              ]}
              color="from-indigo-500 to-indigo-600"
              link="/integrations"
            />
          </div>
        </div>
      </section>

      {/* CTA */}
      <section className="bg-gradient-to-r from-blue-600 to-purple-600 py-20">
        <div className="max-w-4xl mx-auto px-4 text-center">
          <h2 className="text-4xl font-bold text-white mb-6">Ready to get started?</h2>
          <Link
            to="/register"
            className="inline-block bg-white text-blue-600 px-8 py-4 rounded-lg font-semibold hover:bg-gray-100 transition shadow-lg"
          >
            Start for free →
          </Link>
        </div>
      </section>
    </div>
  );
}

function FeatureDetail({ icon, title, description, features, color, link }) {
  const content = (
    <div className="bg-gradient-to-br from-gray-50 to-white p-8 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-2">
      <div className={`bg-gradient-to-r ${color} w-16 h-16 rounded-xl flex items-center justify-center mb-6 text-white`}>
        {icon}
      </div>
      <h3 className="text-2xl font-bold text-gray-900 mb-4">{title}</h3>
      <p className="text-gray-600 mb-6 leading-relaxed">{description}</p>
      <ul className="space-y-3">
        {features.map((feature, idx) => (
          <li key={idx} className="flex items-center text-gray-700">
            <CheckCircle className="w-5 h-5 text-green-500 mr-3" />
            {feature}
          </li>
        ))}
      </ul>
      {link && (
        <div className="mt-6 text-blue-600 font-semibold flex items-center">
          Learn more <ArrowRight className="w-4 h-4 ml-2" />
        </div>
      )}
    </div>
  );

  return link ? <Link to={link}>{content}</Link> : <div>{content}</div>;
}

export default FeaturesPage;