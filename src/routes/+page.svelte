<script>
  import CalendarIcon from '../components/icons/Calendar.svelte';
  import CryptioLogo from '../components/icons/Cryptio.svelte';
  import EducationIcon from '../components/icons/Education.svelte';
  import ExportJapanLogo from '../components/icons/ExportJapan.svelte';
  import HeekLogo from '../components/icons/Heek.svelte';
  import JobIcon from '../components/icons/Job.svelte';
  import MapIcon from '../components/icons/Map.svelte';
  import PayfitLogo from '../components/icons/Payfit.svelte';

  const socials = [
    {
      href: 'mailto:melkir13@gmail.com',
      icon: 'bx bx-envelope'
    },
    {
      href: 'https://github.com/melkir',
      icon: 'bx bxl-github'
    },
    {
      href: 'https://line.me/ti/p/VEfTVnxksB',
      icon: 'bx bx-message'
    },
    {
      href: 'https://www.linkedin.com/in/thibault-vieux',
      icon: 'bx bxl-linkedin'
    }
  ];

  const skills = ['Rust', 'TypeScript', 'React', 'GraphQL', 'Docker', 'SQL', 'Git', 'MongoDB'];

  const projects = [
    { href: 'https://github.com/melkir/next-remote-gpio', name: 'Remote GPIO' },
    { href: 'https://github.com/melkir/hacking-workshop', name: 'Pentest Workshop' },
    { href: 'https://github.com/melkir/android-sms-push', name: 'SMS Push' },
    { href: 'https://github.com/melkir/API-Libraries-Demo', name: 'API Libraries Demo' }
  ];

  const experience = [
    {
      title: 'Senior Software Engineer',
      company: 'Cryptio',
      location: 'Paris',
      date: 'April - July 2023',
      type: 'Full-time',
      duration: '4 mos',
      description:
        'Transform digital asset data into auditable records for accounting, treasury, and tax fillings.',
      logo: CryptioLogo
    },
    {
      title: 'Fullstack Software Engineer',
      company: 'PayFit',
      location: 'Paris',
      date: 'Apr 2018 – Mar 2023',
      type: 'Full-time',
      duration: '5 yrs',
      description:
        "Development of the company's internal tools. Including the automation of recurring tasks with government authorities as well as the management of our clients data through intuitive dashboards.",
      logo: PayfitLogo
    },
    {
      title: 'Fullstack Developer',
      company: 'Heek',
      location: 'Paris',
      date: 'Apr – Oct 2017',
      type: 'Full-time',
      duration: '6 mos',
      description:
        'Creation of an intelligent chatbot capable of generating websites according to customers needs.',
      logo: HeekLogo
    },
    {
      title: 'Mobile Application Developer',
      company: 'Export Japan',
      location: 'Osaka',
      date: 'May – Sep 2016',
      type: 'Internship',
      duration: '4 mos',
      description:
        'Built a hybrid mobile app capable of taking advantage of the smartphone hardware capabilities.',
      logo: ExportJapanLogo
    }
  ];

  const education = [
    {
      title: 'Master in Computer Science',
      domain: 'Smart Systems',
      school: 'CY Tech',
      period: '2015 – 2017',
      location: 'Paris',
      image: 'cytech.jpg'
    },
    {
      title: 'B.E. in Computer Science',
      domain: 'Artificial Intelligence',
      school: 'Univ. Paris Cité',
      period: '2013 – 2015',
      location: 'Paris',
      image: 'paris_universite.jpg'
    }
  ];
</script>

<div class="flex items-center sub-page">
  <div class="grid gap-4 grid-cols-3">
    <div class="space-y-4">
      <!-- Start Left Side -->
      <div class="shadow rounded-xl overflow-hidden">
        <!-- Start User Block -->
        <div class="h-24 bg-cover" style="background-image: url('cover.jpg')" />
        <div class="pt-14 p-7 bg-white relative">
          <img src="avatar.jpg" alt="Avatar" class="user-photo" />
          <div class="text-lg font-semibold mb-1.5">Thibault Vieux</div>
          <div class="text-sm text-gray-400 mb-1.5">Software Engineer</div>
          <ul class="flex space-x-4">
            {#each socials as { href, icon }}
              <li>
                <a {href} target="_blank" class="social-link-hover">
                  <i class={`bx ${icon} text-xl`} />
                </a>
              </li>
            {/each}
          </ul>
        </div>
      </div>
      <!-- End User Block -->
      <div class="p-5 block-section">
        <!-- Start Info Block -->
        <h2 class="block-title">Information</h2>
        <div class="space-y-4">
          <div class="flex justify-between">
            <div class="text-gray-400">Location</div>
            <div class="font-medium text-right text-gray-600">France</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">Experience</div>
            <div class="font-medium text-right text-gray-600">6 years</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">Available in</div>
            <div class="font-medium text-right text-gray-600">1 week</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">Relocation</div>
            <div class="font-medium text-right text-gray-600">Amenable</div>
          </div>
          <div class="flex justify-between">
            <div class="text-gray-400">Language</div>
            <div class="flex content-center">
              <span class="bg-gray-100 text-gray-800 font-medium text-xs mr-1 px-1 py-1 rounded">
                FR 🇫🇷
              </span>
              <span class="bg-gray-100 text-gray-800 font-medium text-xs px-1 py-1 rounded">
                EN 🇬🇧
              </span>
            </div>
          </div>
        </div>
      </div>
      <!-- End Info Block -->
      <div class="p-5 block-section flow-root">
        <!-- Start Skills Block -->
        <h2 class="block-title">Skills</h2>
        <div class="-m-2 flex flex-wrap">
          {#each skills as skill}
            <span class="skill-tag bg-purple-100 text-purple-800">{skill}</span>
          {/each}
        </div>
      </div>
      <!-- End Skills Block -->
      <div class="p-5 block-section flow-root">
        <!-- Start Projects Block -->
        <h2 class="block-title">Key Projects</h2>
        <div class="-m-2 flex flex-wrap">
          {#each projects as { href, name }}
            <a {href} target="_blank" class="skill-tag bg-blue-100 text-blue-800">{name}</a>
          {/each}
        </div>
      </div>
      <!-- End Projects Block -->
    </div>
    <!-- End Left Side -->

    <div class="space-y-3 col-span-2">
      <!-- Start Right Side -->
      <div class="p-5 block-section">
        <!-- Start About Me Block -->
        <h2 class="block-title">About me</h2>
        <p class="text-gray-600 text-sm">
          I specialize in <b>building web applications</b> and currently simplify digital asset
          accounting at Cryptio. In my spare time, I program in Rust and work on home automation
          projects. I also taught <b>penetration testing</b> and
          <b>network security</b> at university, gaining valuable experience and feedback.
        </p>
      </div>
      <!-- End About Me Block -->
      <div class="p-5 block-section">
        <!-- Start Experience Block -->
        <h2 class="block-title">Experience</h2>
        {#each experience as { company, date, description, duration, location, logo, title, type }, i}
          <div class={i === experience.length - 1 ? 'item-section' : 'mb-4 item-section'}>
            <svelte:component this={logo} />
            <div class="w-full space-y-4">
              <div class="item-header">
                <div class="space-y-1.5">
                  <div class="font-medium">{title}</div>
                  <div class="flex space-x-4">
                    <div class="item-header-info">
                      <JobIcon />
                      <span>{company}</span>
                    </div>
                    <div class="item-header-info">
                      <MapIcon />
                      <span>{location}</span>
                    </div>
                  </div>
                </div>
                <div class="space-y-2 text-right">
                  <div class="job-item-badge">{type} · {duration}</div>
                  <div class="item-header-info">
                    <CalendarIcon />
                    <span>{date}</span>
                  </div>
                </div>
              </div>
              <p class="text-gray-600 text-sm" style="margin-top: 0.5rem">
                {description}
              </p>
              {#if i !== experience.length - 1}
                <div class="border-b border-gray-200" />
              {/if}
            </div>
          </div>
        {/each}
      </div>
      <!-- End Experience Block -->
      <div class="p-5 block-section">
        <!-- Start Education Block -->
        <h2 class="block-title">Education</h2>
        {#each education as { domain, image, location, period, school, title }, i}
          <div class={i === education.length - 1 ? 'item-section' : 'mb-4 item-section'}>
            <img alt="logo" class="rounded-md" src={image} height="48px" width="48px" />
            <div class="w-full space-y-4">
              <div class="item-header">
                <div class="space-y-1.5">
                  <div class="font-medium">{title}</div>
                  <div class="flex space-x-4">
                    <div class="item-header-info">
                      <EducationIcon />
                      <span>{school}</span>
                    </div>
                    <div class="item-header-info">
                      <MapIcon />
                      <span>{location}</span>
                    </div>
                  </div>
                </div>
                <div class="space-y-2 text-right">
                  <div class="job-item-badge">{domain}</div>
                  <div class="item-header-info">
                    <CalendarIcon />
                    <span>{period}</span>
                  </div>
                </div>
              </div>
              {#if i !== education.length - 1}
                <div class="border-b border-gray-200" />
              {/if}
            </div>
          </div>
        {/each}
      </div>
      <!-- End Education Block -->
    </div>
    <!-- End Right Side -->
  </div>
</div>
