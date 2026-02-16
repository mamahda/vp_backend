insert into roles (name) values 
('admin'),
('agent'),
('customer');

INSERT INTO property_types (id, name) VALUES
(1, 'Rumah'),
(2, 'Apartment'),
(3, 'Ruko'),
(4, 'Tanah');

INSERT INTO users (email, password, username, phone_number, role_id) 
VALUES 
('budi.agent@victoria.com', '$2a$10$Q4x/ZViUI0hQJLuVG7v1ZudEVPlImaIqyWf3bhYBZWAStjlyj6rUK', 'Budi Santoso', '081234567890', 2),
('siti.agent@victoria.com', '$2a$10$Q4x/ZViUI0hQJLuVG7v1ZudEVPlImaIqyWf3bhYBZWAStjlyj6rUK', 'Siti Aminah', '081298765432', 2),
('andi.agent@victoria.com', '$2a$10$Q4x/ZViUI0hQJLuVG7v1ZudEVPlImaIqyWf3bhYBZWAStjlyj6rUK', 'Andi Wijaya', '081311223344', 2);

INSERT INTO properties (
    title,
    description,
    price,
    status,
    province,
    regency,
    district,
    address,
    building_area,
    land_area,
    electricity,
    water_source,
    bedrooms,
    bathrooms,
    floors,
    garage,
    carport,
    certificate,
    year_constructed,
    sale_type,
    latitude,
    longitude,
    created_at,
    cover_image_url,
    property_type_id,
    agent_id
) VALUES
-- 1-3: Surabaya (Agent 1)
(
    'Rumah Minimalis Modern 2 Lantai',
    'Rumah minimalis modern di kawasan strategis, dekat sekolah dan pusat perbelanjaan.',
    850000000, 1, 'Jawa Timur', 'Surabaya', 'Sukolilo', 'Jl. Keputih No. 15',
    120, 150, 2200, 1, 3, 2, 2, 1, 1, 'SHM', 2018, 'jual',
    -7.2858, 112.7964, NOW() - INTERVAL 5 DAY, 'uploads/cover-1.jpg', 1, 1
),
(
    'Townhouse di Citraland Surabaya',
    'Townhouse mewah di kawasan Citraland, fasilitas lengkap: clubhouse, kolam renang.',
    1200000000, 1, 'Jawa Timur', 'Surabaya', 'Pakuwon', 'Cluster Jasmine No. 8',
    180, 200, 4400, 1, 4, 3, 2, 2, 1, 'SHM', 2020, 'jual',
    -7.2892, 112.6454, NOW() - INTERVAL 3 DAY, 'uploads/cover-2.jpg', 1, 1
),
(
    'Ruko Strategis di Jalan Protokol',
    'Ruko 2 lantai di lokasi sangat strategis, depan jalan protokol, ramai pengunjung.',
    2500000000, 1, 'Jawa Timur', 'Surabaya', 'Genteng', 'Jl. Basuki Rahmat No. 45',
    160, 120, 6600, 2, 0, 2, 2, 0, 0, 'HGB', 2015, 'jual',
    -7.2638, 112.7402, NOW() - INTERVAL 1 DAY, 'uploads/cover-3.jpg', 3, 1
),

-- 4-6: Bandung (Agent 2)
(
    'Rumah Tipe 36 di Bandung Utara',
    'Rumah sederhana di Bandung Utara, udara sejuk, lingkungan asri.',
    350000000, 1, 'Jawa Barat', 'Bandung', 'Cimenyan', 'Jl. Bukit Pakar Timur No. 10',
    36, 60, 1300, 1, 2, 1, 1, 0, 1, 'SHM', 2010, 'jual',
    -6.8667, 107.6333, NOW() - INTERVAL 7 DAY, 'uploads/cover-4.jpg', 1, 2
),
(
    'Apartemen Studio di Setiabudi',
    'Apartemen studio dengan view kota Bandung, fully furnished.',
    450000000, 1, 'Jawa Barat', 'Bandung', 'Setiabudi', 'Apartemen Green Valley Tower B Lantai 12',
    28, 28, 900, 2, 1, 1, 12, 1, 0, 'Strata Title', 2019, 'jual',
    -6.8604, 107.5921, NOW() - INTERVAL 2 DAY, 'uploads/cover-5.jpg', 2, 2
),
(
    'Villa di Lembang untuk Staycation',
    'Villa mewah dengan kolam renang pribadi, 4 kamar tidur, view gunung.',
    3500000000, 1, 'Jawa Barat', 'Bandung', 'Lembang', 'Jl. Maribaya No. 88',
    300, 800, 7700, 3, 4, 4, 2, 3, 2, 'SHM', 2021, 'jual',
    -6.8222, 107.6147, NOW(), 'uploads/cover-6.jpg', 1, 2
),

-- 7-9: Jakarta (Agent 3)
(
    'Apartemen Lux di SCBD Sudirman',
    'Apartemen premium di kawasan SCBD, 2BR, fasilitas lengkap: gym, pool.',
    3200000000, 1, 'DKI Jakarta', 'Jakarta Selatan', 'Setiabudi', 'Pacific Place Residence Tower A',
    85, 85, 4400, 2, 2, 2, 35, 2, 0, 'Strata Title', 2022, 'jual',
    -6.2241, 106.8097, NOW() - INTERVAL 4 DAY, 'uploads/cover-7.jpg', 2, 3
),
(
    'Kantor di Kuningan Jakarta',
    'Office space di gedung premium Kuningan, 1 lantai penuh, siap pakai.',
    8500000000, 1, 'DKI Jakarta', 'Jakarta Selatan', 'Kuningan', 'Jl. HR Rasuna Said Kav. C-1',
    300, 300, 35000, 2, 0, 4, 10, 10, 0, 'HGB', 2018, 'jual',
    -6.2247, 106.8326, NOW() - INTERVAL 6 DAY, 'uploads/cover-8.jpg', 3, 3
),
(
    'Rumah Mewah di Pondok Indah',
    'Rumah mansion di cluster eksklusif Pondok Indah, 6 kamar, kolam renang.',
    28000000000, 1, 'DKI Jakarta', 'Jakarta Selatan', 'Kebayoran Lama', 'Cluster Pondok Indah Residence',
    800, 1200, 22000, 1, 6, 5, 3, 4, 3, 'SHM', 2020, 'jual',
    -6.2694, 106.7821, NOW() - INTERVAL 9 DAY, 'uploads/cover-9.jpg', 1, 3
),

-- 10-12: Bali (Agent 1)
(
    'Villa di Seminyak Bali',
    'Villa private dengan traditional Balinese style, 3BR, dekat pantai.',
    4500000000, 1, 'Bali', 'Badung', 'Kuta', 'Jl. Kayu Aya No. 33',
    250, 400, 5500, 3, 3, 3, 2, 2, 1, 'Hak Pakai', 2019, 'jual',
    -8.6853, 115.1557, NOW() - INTERVAL 8 DAY, 'uploads/cover-10.jpg', 1, 1
),
(
    'Tanah Kavling di Canggu',
    'Tanah siap bangun di area Canggu yang sedang berkembang pesat.',
    1800000000, 1, 'Bali', 'Badung', 'North Kuta', 'Jl. Pantai Berawa',
    0, 300, 0, 0, 0, 0, 0, 0, 0, 'SHM', 0, 'jual',
    -8.6644, 115.1418, NOW() - INTERVAL 10 DAY, 'uploads/cover-11.jpg', 4, 1
),
(
    'Guest House di Ubud',
    'Guest house beroperasi di Ubud, 8 kamar, average occupancy 70%.',
    3500000000, 1, 'Bali', 'Gianyar', 'Ubud', 'Jl. Monkey Forest',
    350, 600, 9900, 2, 8, 8, 2, 5, 3, 'Hak Pakai', 2017, 'jual',
    -8.5069, 115.2625, NOW() - INTERVAL 12 DAY, 'uploads/cover-12.jpg', 1, 1
),

-- 13-15: Yogyakarta (Agent 2)
(
    'Rumah Murah di Jogja Utara',
    'Rumah ekonomis di Jogja, dekat kampus UGM, cocok untuk keluarga muda.',
    250000000, 1, 'DI Yogyakarta', 'Sleman', 'Depok', 'Jl. Kaliurang Km 5.5',
    45, 72, 900, 1, 2, 1, 1, 0, 1, 'SHM', 2012, 'jual',
    -7.7589, 110.3785, NOW() - INTERVAL 15 DAY, 'uploads/cover-13.jpg', 1, 2
),
(
    'Ruko di Malioboro',
    'Ruko 3 lantai di kawasan Malioboro, strategis untuk retail.',
    4500000000, 1, 'DI Yogyakarta', 'Yogyakarta', 'Gondomanan', 'Jl. Malioboro No. 123',
    180, 150, 7700, 2, 0, 3, 3, 0, 0, 'HGB', 2014, 'jual',
    -7.7944, 110.3659, NOW() - INTERVAL 14 DAY, 'uploads/cover-14.jpg', 3, 2
),
(
    'Kost Eksklusif di Jogja',
    'Kost wanita 12 kamar, full AC, ada wifi, dapur bersama.',
    1200000000, 1, 'DI Yogyakarta', 'Yogyakarta', 'Umbulharjo', 'Jl. C. Simanjuntak No. 15',
    200, 300, 8800, 1, 12, 6, 3, 3, 2, 'SHM', 2016, 'jual',
    -7.8163, 110.3862, NOW() - INTERVAL 13 DAY, 'uploads/cover-15.jpg', 1, 2
),

-- 16-18: Medan (Agent 3)
(
    'Rumah Taman di Medan Selatan',
    'Rumah dengan taman luas di Medan, lingkungan elite, keamanan 24 jam.',
    950000000, 1, 'Sumatera Utara', 'Medan', 'Medan Selatan', 'Jl. Pattimura No. 88',
    220, 350, 4400, 1, 4, 3, 2, 2, 1, 'SHM', 2015, 'jual',
    3.5719, 98.6601, NOW() - INTERVAL 18 DAY, 'uploads/cover-16.jpg', 1, 3
),
(
    'Ruko di Jalan Gatot Subroto',
    'Ruko baru 2 lantai di pusat bisnis Medan. Cocok untuk showroom.',
    2200000000, 1, 'Sumatera Utara', 'Medan', 'Medan Petisah', 'Jl. Gatot Subroto No. 45',
    160, 140, 5500, 2, 0, 2, 2, 0, 0, 'HGB', 2021, 'jual',
    3.5891, 98.6517, NOW() - INTERVAL 17 DAY, 'uploads/cover-17.jpg', 3, 3
),
(
    'Apartemen di Centre Point',
    'Apartemen 1BR di gedung baru, view kota Medan.',
    650000000, 1, 'Sumatera Utara', 'Medan', 'Medan Barat', 'Centre Point Apartment Tower B',
    48, 48, 2200, 2, 1, 1, 25, 1, 0, 'Strata Title', 2023, 'jual',
    3.5917, 98.6811, NOW() - INTERVAL 16 DAY, 'uploads/cover-18.jpg', 2, 3
),

-- 19-21: Makassar (Agent 1)
(
    'Rumah di Pantai Losari',
    'Rumah dengan view langsung ke pantai Losari. Lokasi premium.',
    2800000000, 1, 'Sulawesi Selatan', 'Makassar', 'Ujung Pandang', 'Jl. Penghibur No. 17',
    180, 250, 6600, 1, 3, 3, 2, 2, 1, 'SHM', 2019, 'jual',
    -5.1444, 119.4061, NOW() - INTERVAL 20 DAY, 'uploads/cover-19.jpg', 1, 1
),
(
    'Kavling di BTP Makassar',
    'Tanah kavling siap bangun di BTP, kawasan terencana.',
    750000000, 1, 'Sulawesi Selatan', 'Makassar', 'Tamalanrea', 'Blok D No. 12',
    0, 200, 0, 0, 0, 0, 0, 0, 0, 'SHM', 0, 'jual',
    -5.1189, 119.5022, NOW() - INTERVAL 19 DAY, 'uploads/cover-20.jpg', 4, 1
),
(
    'Ruko di Panakkukang',
    'Ruko di kawasan komersial Panakkukang, strategis untuk usaha.',
    1850000000, 1, 'Sulawesi Selatan', 'Makassar', 'Panakkukang', 'Jl. Metro Tj. Bunga',
    140, 120, 4400, 2, 2, 2, 2, 0, 0, 'HGB', 2018, 'jual',
    -5.1558, 119.4475, NOW() - INTERVAL 21 DAY, 'uploads/cover-21.jpg', 3, 1
),

-- 22-24: Sewa (Mixed Agents)
(
    'Kontrakan Bulanan di Surabaya',
    'Kontrakan bulanan di Surabaya Pusat, dekat kantor dan mall.',
    3500000, 2, 'Jawa Timur', 'Surabaya', 'Tegalsari', 'Jl. Kedungdoro No. 22',
    30, 30, 900, 2, 1, 1, 1, 0, 0, 'SHM', 2015, 'sewa',
    -7.2611, 112.7345, NOW() - INTERVAL 22 DAY, 'uploads/cover-22.jpg', 1, 1
),
(
    'Apartemen Harian di Bandung',
    'Apartemen harian untuk staycation, fully furnished.',
    450000, 2, 'Jawa Barat', 'Bandung', 'Coblong', 'Jl. Siliwangi No. 8',
    32, 32, 1300, 2, 1, 1, 8, 1, 0, 'Strata Title', 2020, 'sewa',
    -6.8833, 107.6101, NOW() - INTERVAL 23 DAY, 'uploads/cover-23.jpg', 2, 2
),
(
    'Kantor Sewa di SCBD',
    'Office space sewa di SCBD, termasuk cleaning service.',
    25000000, 2, 'DKI Jakarta', 'Jakarta Selatan', 'Setiabudi', 'Sudirman Central Business District',
    50, 50, 5500, 2, 0, 1, 12, 5, 0, 'HGB', 2021, 'sewa',
    -6.2239, 106.8115, NOW() - INTERVAL 24 DAY, 'uploads/cover-24.jpg', 3, 3
),

-- 25-27: Properti Murah (Agent 2)
(
    'Rumah Subsidi di Bekasi',
    'Rumah tipe 36 program pemerintah, cocok untuk keluarga pertama.',
    180000000, 1, 'Jawa Barat', 'Bekasi', 'Bekasi Timur', 'Cluster Griya Asri Blok A1',
    36, 60, 900, 1, 2, 1, 1, 0, 1, 'SHM', 2022, 'jual',
    -6.2541, 107.0182, NOW() - INTERVAL 25 DAY, 'uploads/cover-25.jpg', 1, 2
),
(
    'Rumah Kos di Malang',
    'Rumah kos 8 kamar di dekat kampus Brawijaya.',
    275000000, 1, 'Jawa Timur', 'Malang', 'Lowokwaru', 'Jl. Veteran No. 15',
    100, 150, 3300, 1, 8, 4, 2, 0, 2, 'SHM', 2014, 'jual',
    -7.9523, 112.6134, NOW() - INTERVAL 26 DAY, 'uploads/cover-26.jpg', 1, 2
),
(
    'Rumah Warisan di Solo',
    'Rumah lawas di Solo, butuh renovasi. Lokasi strategis.',
    220000000, 1, 'Jawa Tengah', 'Surakarta', 'Laweyan', 'Jl. Slamet Riyadi No. 150',
    80, 120, 1300, 1, 3, 2, 1, 0, 1, 'SHM', 1975, 'jual',
    -7.5684, 110.8214, NOW() - INTERVAL 27 DAY, 'uploads/cover-27.jpg', 1, 2
),

-- 28-30: Properti Mewah (Agent 3)
(
    'Penthouse di Plaza Indonesia',
    'Penthouse mewah di atas mall Plaza Indonesia, 3BR, private elevator.',
    45000000000, 1, 'DKI Jakarta', 'Jakarta Pusat', 'Menteng', 'Plaza Indonesia Residence',
    350, 350, 22000, 2, 3, 4, 48, 3, 0, 'Strata Title', 2023, 'jual',
    -6.1919, 106.8227, NOW() - INTERVAL 28 DAY, 'uploads/cover-28.jpg', 2, 3
),
(
    'Hotel Bintang 4 di Bali',
    'Hotel beroperasi di Kuta, 40 kamar, average occupancy 75%.',
    85000000000, 1, 'Bali', 'Badung', 'Kuta', 'Jl. Legian No. 99',
    2000, 3000, 66000, 3, 40, 45, 6, 30, 15, 'Hak Pakai', 2018, 'jual',
    -8.7111, 115.1741, NOW() - INTERVAL 29 DAY, 'uploads/cover-29.jpg', 1, 3
),
(
    'Resort di Labuan Bajo',
    'Resort dengan 15 villa private view laut, bisnis berjalan baik.',
    120000000000, 1, 'Nusa Tenggara Timur', 'Manggarai Barat', 'Komodo', 'Pulau Seraya',
    1200, 5000, 33000, 3, 15, 20, 2, 20, 10, 'Hak Pakai', 2021, 'jual',
    -8.4354, 119.8252, NOW() - INTERVAL 30 DAY, 'uploads/cover-30.jpg', 1, 3
);

insert into users (email, password, username, phone_number, role_id) values 
('admin@admin.com', '$2a$10$Q4x/ZViUI0hQJLuVG7v1ZudEVPlImaIqyWf3bhYBZWAStjlyj6rUK', 'adminuser', '1234567890', 1);
